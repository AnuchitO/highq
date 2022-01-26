package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/anuchito/dbstore/pb"
	"github.com/anuchito/dbstore/pkg/version"
)

// NewMainHandler creates all http handlers
func NewMainHandler(filename string) http.Handler {
	r := http.NewServeMux()
	mydb := NewDb(filename)
	service := NewService(mydb)
	myhandler := NewHandler(service)
	r.Handle("/db/", ErrorMiddleware(myhandler.HandleDb))
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Handle("/version", ErrorMiddleware(versionHandler))
	return r
}

func versionHandler(w http.ResponseWriter, r *http.Request) error {
	info := struct {
		BuildTime string `json:"buildTime"`
		Commit    string `json:"commit"`
		Release   string `json:"release"`
	}{
		version.BuildTime, version.Commit, version.Release,
	}
	body, err := json.Marshal(info)
	if err != nil {
		return fmt.Errorf("Could not encode version data: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		return fmt.Errorf("Error writing to http response: %v", err)
	}
	return nil
}

// Handler holds all http methods
type Handler struct {
	service *Service
}

// HTTPError contains http status code
type HTTPError struct {
	Err        error
	StatusCode int
	Message    string
}

// NewHTTPError constructor
func NewHTTPError(err error, statusCode int, message string) *HTTPError {
	return &HTTPError{Err: err, StatusCode: statusCode, Message: message}
}

func (e *HTTPError) Error() string {
	// only log detailed error, don't return it to client
	fmt.Printf(e.Err.Error() + ": " + e.Message + "\n")
	return e.Message
}

// ErrorMiddleware wraps a normal handler and converts errors to corresponding http status codes
type ErrorMiddleware func(http.ResponseWriter, *http.Request) error

// ErrorJSON http response type
type ErrorJSON struct {
	Message string `json:"message"`
}

func (fn ErrorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err != nil {
		switch err.(type) {
		case *HTTPError:
			e := err.(*HTTPError)
			info := &ErrorJSON{
				Message: e.Error(),
			}
			body, err := json.Marshal(info)
			if err != nil {
				fmt.Printf("Could not encode error data: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(e.StatusCode)
			_, err = w.Write(body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		default:
			info := &ErrorJSON{
				Message: err.Error(),
			}
			body, err := json.Marshal(info)
			if err != nil {
				fmt.Printf("Could not encode error data: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			_, err = w.Write(body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}

// NewHandler is a constructor of all handlers
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func getID(s string) (string, error) {
	id := strings.TrimPrefix(s, "/db/")
	arr := strings.Split(id, "/")
	if len(arr) != 1 {
		return "", fmt.Errorf("id contains a slash /, %s", id)
	}
	return id, nil
}

// HandleDb handles all http routes
func (h *Handler) HandleDb(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return h.getHandler(w, r)
	} else if r.Method == "POST" {
		return h.setHandler(w, r)
	} else if r.Method == "DELETE" {
		return h.deleteHandler(w, r)
	}
	return NewHTTPError(fmt.Errorf(""), http.StatusNotFound, "NOT FOUND")
}

func (h *Handler) setHandler(w http.ResponseWriter, r *http.Request) error {
	key, err := getID(r.URL.Path)
	if err != nil {
		return NewHTTPError(err, http.StatusBadRequest, "requested key not valid")
	}

	if r.Header.Get("Content-Type") != "application/octet-stream" {
		return NewHTTPError(fmt.Errorf(""), http.StatusBadRequest, "Mime-Type not supported, application/octet-stream is supported")
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	entity := &pb.Entity{Key: key, Value: body}
	err = h.service.Set(entity)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	return nil
}

func (h *Handler) getHandler(w http.ResponseWriter, r *http.Request) error {
	key, err := getID(r.URL.Path)
	if err != nil {
		return NewHTTPError(err, http.StatusBadRequest, "requested key not valid")
	}
	entity, err := h.service.Get(key)
	if err != nil {
		return NewHTTPError(err, http.StatusInternalServerError, "error GET the requested key")
	}
	if entity == nil {
		return NewHTTPError(fmt.Errorf(""), http.StatusNotFound, "key does not exist")
	}

	if r.URL.Query().Get("format") == "json" {
		w.Header().Set("Content-Type", "application/json")
	}
	fmt.Fprintf(w, "%s", entity.Value)
	return nil
}

func (h *Handler) deleteHandler(w http.ResponseWriter, r *http.Request) error {
	key, err := getID(r.URL.Path)
	if err != nil {
		return NewHTTPError(err, http.StatusBadRequest, "requested key not valid")
	}
	err = h.service.Delete(key)
	if err != nil {
		return NewHTTPError(err, http.StatusBadRequest, "requested key does not exist")
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

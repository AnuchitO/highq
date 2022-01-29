package http

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestMakeHTTPCall(t *testing.T) {
	testTable := []struct {
		name             string
		server           *httptest.Server
		expectedResponse *Response
		expectedErr      error
	}{
		{
			name: "happy-server-response",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"id": 1, "name": "kyle", "description": "novice gopher"}`))
			})),
			expectedResponse: &Response{
				ID:          1,
				Name:        "kyle",
				Description: "novice gopher",
			},
			expectedErr: nil,
		},
	}
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			defer tc.server.Close()
			resp, err := MakeHTTPCall(tc.server.URL)
			if !reflect.DeepEqual(resp, tc.expectedResponse) {
				t.Errorf("expected (%v), got (%v)", tc.expectedResponse, resp)
			}
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected (%v), got (%v)", tc.expectedErr, err)
			}
		})
	}
}

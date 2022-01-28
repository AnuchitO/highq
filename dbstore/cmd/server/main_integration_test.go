//go:build integration

package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestStoreSuccess(t *testing.T) {
	p, _ := rand.Prime(rand.Reader, 10)
	key := "key" + p.String()
	val := "val" + p.String()
	url := fmt.Sprintf("http://localhost:%s/db/%s", os.Getenv("PORT"), key)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(val))
	req.Header.Add("Content-Type", "application/octet-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Error("Expected status code to be 201, got", resp.StatusCode)
	}
}

func TestShouldGetJsonData(t *testing.T) {
	p, _ := rand.Prime(rand.Reader, 10)
	key := "key" + p.String()
	val := `{"name":"AnuchitO"}`
	url := fmt.Sprintf("http://localhost:%s/db/%s", os.Getenv("PORT"), key)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(val))
	req.Header.Add("Content-Type", "application/octet-stream")

	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusCreated {
		t.Error("Expected status code to be 201, got", r.StatusCode)
	}

	q := `format=json`
	baseURL := fmt.Sprintf("http://localhost:%s/db/%s?%s", os.Getenv("PORT"), key, q)

	res, err := http.Get(baseURL)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Error("Expected status code to be 200, got", res.StatusCode)
	}
	if res.Header.Get("content-type") != "application/json" {
		t.Errorf("Expected content-type to be application/json, got %#v\n", res.Header)
	}
}

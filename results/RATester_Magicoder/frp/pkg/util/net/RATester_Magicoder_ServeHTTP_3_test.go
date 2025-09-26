package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := &HTTPGzipWrapper{
		h: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello, world!"))
		}),
	}

	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Header.Get("Content-Encoding") != "gzip" {
		t.Error("Expected Content-Encoding: gzip")
	}

	if string(body) != "Hello, world!" {
		t.Error("Expected body to be 'Hello, world!'")
	}
}

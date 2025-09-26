package common

import (
	"fmt"
	"net/http"
	"testing"
)

func TestChangeHostAndHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &http.Request{}
	ChangeHostAndHeader(r, "host", "header", "addr", true)
	if r.Host != "host" {
		t.Errorf("r.Host = %v, want %v", r.Host, "host")
	}
	if r.Header.Get("header") != "header" {
		t.Errorf("r.Header.Get(\"header\") = %v, want %v", r.Header.Get("header"), "header")
	}
	if r.Header.Get("X-Forwarded-For") != "addr" {
		t.Errorf("r.Header.Get(\"X-Forwarded-For\") = %v, want %v", r.Header.Get("X-Forwarded-For"), "addr")
	}
	if r.Header.Get("X-Real-IP") != "addr" {
		t.Errorf("r.Header.Get(\"X-Real-IP\") = %v, want %v", r.Header.Get("X-Real-IP"), "addr")
	}
}

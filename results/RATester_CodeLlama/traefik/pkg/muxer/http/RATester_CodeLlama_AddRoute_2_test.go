package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAddRoute_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Muxer{}
	rule := "Host(`example.com`) && Path(`/foo`) && Method(`GET`)"
	syntax := "v2"
	priority := 1
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {})
	err := m.AddRoute(rule, syntax, priority, handler)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

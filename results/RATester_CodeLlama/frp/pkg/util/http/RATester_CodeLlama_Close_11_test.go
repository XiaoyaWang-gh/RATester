package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClose_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Server{}
	s.hs = &http.Server{}
	err := s.Close()
	if err != nil {
		t.Errorf("Close() error = %v, want nil", err)
		return
	}
}

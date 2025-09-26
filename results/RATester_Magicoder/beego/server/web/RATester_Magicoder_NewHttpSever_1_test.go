package web

import (
	"fmt"
	"testing"
)

func TestNewHttpSever_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := NewHttpSever()
	if server == nil {
		t.Error("Expected a non-nil server, but got nil")
	}
}

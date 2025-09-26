package httpserver

import (
	"fmt"
	"testing"
)

func TestBindAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srv := &Server{
		bindAddr: "127.0.0.1",
	}

	expected := "127.0.0.1"
	actual := srv.BindAddr()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

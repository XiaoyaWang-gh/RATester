package streamserver

import (
	"fmt"
	"testing"
)

func TestBindAddr_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &Server{
		bindAddr: "127.0.0.1",
	}

	expected := "127.0.0.1"
	actual := server.BindAddr()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}

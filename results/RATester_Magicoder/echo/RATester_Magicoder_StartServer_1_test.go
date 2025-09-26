package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestStartServer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{}
	s := &http.Server{}

	err := e.StartServer(s)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

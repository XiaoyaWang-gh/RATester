package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestConfigureServer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{}
	s := &http.Server{}

	err := e.configureServer(s)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Add more test cases as needed
}

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

	// setup
	var e Echo
	var s *http.Server
	// test
	if err := e.configureServer(s); err != nil {
		t.Errorf("error = %v", err)
	}
	// teardown
}

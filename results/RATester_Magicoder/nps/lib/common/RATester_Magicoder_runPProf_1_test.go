package common

import (
	"fmt"
	"net/http"
	"testing"
)

func TestrunPProf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ipPort := "localhost:8080"
	runPProf(ipPort)

	// Test if the server is running
	resp, err := http.Get("http://" + ipPort + "/debug/pprof/")
	if err != nil {
		t.Errorf("Failed to connect to server: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

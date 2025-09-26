package common

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRunPProf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ipPort := "localhost:6060"
	runPProf(ipPort)

	resp, err := http.Get(fmt.Sprintf("http://%s/debug/pprof/", ipPort))
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
	}
}

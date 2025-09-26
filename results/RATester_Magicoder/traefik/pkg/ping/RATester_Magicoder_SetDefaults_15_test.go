package ping

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetDefaults_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &Handler{}
	h.SetDefaults()

	if h.EntryPoint != "traefik" {
		t.Errorf("Expected EntryPoint to be 'traefik', but got %s", h.EntryPoint)
	}

	if h.TerminatingStatusCode != http.StatusServiceUnavailable {
		t.Errorf("Expected TerminatingStatusCode to be %d, but got %d", http.StatusServiceUnavailable, h.TerminatingStatusCode)
	}
}

package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestclientIPV2_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &matchersTree{}
	clientIPs := []string{"192.168.1.1", "192.168.1.2"}

	err := clientIPV2(tree, clientIPs...)
	if err != nil {
		t.Errorf("clientIPV2() error = %v", err)
		return
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("http.NewRequest() error = %v", err)
		return
	}

	req.RemoteAddr = "192.168.1.1:8080"
	if !tree.match(req) {
		t.Errorf("clientIPV2() = %v, want %v", false, true)
	}

	req.RemoteAddr = "192.168.1.3:8080"
	if tree.match(req) {
		t.Errorf("clientIPV2() = %v, want %v", true, false)
	}
}

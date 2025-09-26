package service

import (
	"fmt"
	"net/http"
	"testing"
)

func TestisWebSocketUpgrade_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	req := &http.Request{
		Header: http.Header{
			"Connection": []string{"Upgrade"},
			"Upgrade":    []string{"websocket"},
		},
	}

	if !isWebSocketUpgrade(req) {
		t.Error("Expected isWebSocketUpgrade to return true")
	}
}

package echo

import (
	"fmt"
	"testing"

	"golang.org/x/net/http2"
)

func TestStartH2CServer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{}
	h2s := &http2.Server{}
	address := ":8080"

	err := e.StartH2CServer(address, h2s)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

package echo

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClose_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &Echo{}
	e.Server = &http.Server{}
	e.TLSServer = &http.Server{}

	err := e.Close()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

package api

import (
	"fmt"
	"testing"
)

func TestProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := tcpMiddlewareRepresentation{}
	m.Provider = "test"
	if m.provider() != "test" {
		t.Errorf("Expected %s, got %s", "test", m.provider())
	}
}

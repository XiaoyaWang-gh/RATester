package testing

import (
	"fmt"
	"testing"
)

func TestGetPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	port := getPort()
	if port != "8080" {
		t.Errorf("getPort() = %v, want %v", port, "8080")
	}
}

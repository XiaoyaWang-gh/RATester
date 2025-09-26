package mock

import (
	"fmt"
	"testing"
)

func TestStartMock_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stub := StartMock()
	if stub == nil {
		t.Errorf("Expected a non-nil stub, got nil")
	}
}

package mock

import (
	"fmt"
	"testing"
)

func TestnewOrmStub_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stub := newOrmStub()
	if stub == nil {
		t.Error("Expected a non-nil stub, but got nil")
	}
	if len(stub.ms) != 0 {
		t.Errorf("Expected an empty slice, but got %v", stub.ms)
	}
}

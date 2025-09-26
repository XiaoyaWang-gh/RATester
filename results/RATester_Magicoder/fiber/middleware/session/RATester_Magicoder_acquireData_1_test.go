package session

import (
	"fmt"
	"testing"
)

func TestacquireData_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	data := acquireData()
	if data == nil {
		t.Error("Expected data to be not nil")
	}
}

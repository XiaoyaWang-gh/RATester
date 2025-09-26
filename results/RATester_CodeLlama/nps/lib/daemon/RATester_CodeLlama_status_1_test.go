package daemon

import (
	"fmt"
	"testing"
)

func TestStatus_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pidPath := "./pid"
	f := "test"
	if status(f, pidPath) {
		t.Log("status is true")
	} else {
		t.Error("status is false")
	}
}

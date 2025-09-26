package client

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := cancelErr{Err: errors.New("test")}
	if e.Error() != "test" {
		t.Errorf("Error() = %v, want %v", e.Error(), "test")
	}
}

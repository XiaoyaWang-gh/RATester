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

	e := cancelErr{Err: errors.New("test error")}
	expected := "test error"
	result := e.Error()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

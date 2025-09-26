package parse

import (
	"errors"
	"fmt"
	"testing"
)

func Testrecover_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	err := errors.New("test error")
	defer tree.recover(&err)
	panic(err)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

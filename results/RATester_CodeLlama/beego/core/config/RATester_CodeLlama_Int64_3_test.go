package config

import (
	"fmt"
	"testing"
)

func TestInt64_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// setup
	key := "key"
	// mock
	// code under test
	actual, err := Int64(key)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	// assert
	if actual != 0 {
		t.Errorf("actual %v != expected %v", actual, 0)
	}
}

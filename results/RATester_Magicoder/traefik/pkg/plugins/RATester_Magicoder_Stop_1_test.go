package plugins

import (
	"fmt"
	"testing"
)

func TestStop_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pp := _PP{
		WStop: func() error {
			return nil
		},
	}

	err := pp.Stop()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

package file

import (
	"fmt"
	"testing"
)

func TestInit_24(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{}
	err := provider.Init()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

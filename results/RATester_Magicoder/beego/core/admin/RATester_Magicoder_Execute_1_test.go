package admin

import (
	"fmt"
	"testing"
)

func TestExecute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &doNothingCommand{}
	result := d.Execute()

	if result.Status != 404 {
		t.Errorf("Expected status 404, got %d", result.Status)
	}

	if result.Error != CommandNotFound {
		t.Errorf("Expected error CommandNotFound, got %v", result.Error)
	}

	if result.Content != nil {
		t.Errorf("Expected nil content, got %v", result.Content)
	}
}

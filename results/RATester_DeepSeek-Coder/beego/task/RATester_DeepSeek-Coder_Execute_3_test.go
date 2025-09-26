package task

import (
	"fmt"
	"testing"
)

func TestExecute_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &listTaskCommand{}
	result := l.Execute()

	if result.Status != 200 {
		t.Errorf("Expected status 200, got %d", result.Status)
	}

	if result.Error != nil {
		t.Errorf("Expected no error, got %v", result.Error)
	}

	if result.Content == nil {
		t.Errorf("Expected non-nil content")
	}

	content, ok := result.Content.([][]string)
	if !ok {
		t.Errorf("Expected content to be of type [][]string, got %T", result.Content)
	}

	for _, row := range content {
		if len(row) != 4 {
			t.Errorf("Expected each row to have 4 elements, got %d", len(row))
		}
	}
}

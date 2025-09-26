package logs

import (
	"fmt"
	"testing"
)

func TestInit_35(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &fileLogWriter{
		Filename: "test.log",
	}

	err := w.Init("")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	err = w.Init(`{"filename": "test.log"}`)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = w.Init(`{"filename": ""}`)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

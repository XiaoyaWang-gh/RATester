package config

import (
	"fmt"
	"testing"
)

func TestInt_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &fakeConfigContainer{
		data: map[string]string{
			"key": "123",
		},
	}

	val, err := fc.Int("key")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != 123 {
		t.Errorf("expected 123, got %d", val)
	}

	_, err = fc.Int("not_exist")
	if err == nil {
		t.Error("expected error, got nil")
	}
}

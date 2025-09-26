package cache

import (
	"fmt"
	"testing"
)

func TestRegister_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Register panicked: %v", r)
		}
	}()

	Register("test", func() Cache {
		return NewMemoryCache()
	})

	if _, ok := adapters["test"]; !ok {
		t.Errorf("Register failed, adapter not found")
	}
}

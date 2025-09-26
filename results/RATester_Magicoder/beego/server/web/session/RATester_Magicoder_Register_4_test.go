package session

import (
	"fmt"
	"testing"
)

func TestRegister_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Register() panicked with: %v", r)
		}
	}()

	Register("test", nil)
}

package logs

import (
	"fmt"
	"testing"
)

func TestRegister_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Register should not panic, but it did panic with: %v", r)
		}
	}()

	Register("test", func() Logger { return nil })
}

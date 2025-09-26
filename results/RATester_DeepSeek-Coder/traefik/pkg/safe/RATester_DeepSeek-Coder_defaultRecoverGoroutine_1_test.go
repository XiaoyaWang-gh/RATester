package safe

import (
	"fmt"
	"testing"
)

func TestDefaultRecoverGoroutine_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("defaultRecoverGoroutine() panicked with error: %v", r)
		}
	}()

	defaultRecoverGoroutine("test error")
}

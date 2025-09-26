package log

import (
	"fmt"
	"testing"
)

func TestPanicw_1(t *testing.T) {
	t.Run("TestPanicw", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		Panicw("Test panic", "key", "value")
	})
}

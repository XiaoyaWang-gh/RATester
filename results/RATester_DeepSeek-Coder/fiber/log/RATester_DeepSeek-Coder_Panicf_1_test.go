package log

import (
	"fmt"
	"testing"
)

func TestPanicf_1(t *testing.T) {
	t.Run("Panicf", func(t *testing.T) {

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
		Panicf("test")
	})
}

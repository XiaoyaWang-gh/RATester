package log

import (
	"fmt"
	"testing"
)

func TestWarnf_1(t *testing.T) {
	t.Run("test case 1:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		format := "test case 1"
		v := []any{1, 2, 3}
		Warnf(format, v...)
	})
	t.Run("test case 2:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		format := "test case 2"
		v := []any{1, 2, 3}
		Warnf(format, v...)
	})
}

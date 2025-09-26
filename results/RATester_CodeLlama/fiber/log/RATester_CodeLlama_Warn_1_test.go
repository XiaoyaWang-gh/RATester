package log

import (
	"fmt"
	"testing"
)

func TestWarn_1(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// given
		v := []any{1, 2, 3}
		// when
		Warn(v...)
		// then
		// TODO: add assertions
	})
}

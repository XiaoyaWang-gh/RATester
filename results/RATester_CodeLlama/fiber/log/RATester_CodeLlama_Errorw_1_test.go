package log

import (
	"fmt"
	"testing"
)

func TestErrorw_1(t *testing.T) {
	t.Parallel()
	t.Run("test case 1", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		// given
		msg := "msg"
		keysAndValues := []any{"key1", "value1", "key2", "value2"}
		// when
		Errorw(msg, keysAndValues...)
		// then
		// TODO: add assertions
	})
}

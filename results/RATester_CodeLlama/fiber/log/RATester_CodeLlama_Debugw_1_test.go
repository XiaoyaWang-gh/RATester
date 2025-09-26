package log

import (
	"fmt"
	"testing"
)

func TestDebugw_1(t *testing.T) {
	t.Run("TestDebugw", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		msg := "msg"
		keysAndValues := []any{"key", "value"}
		Debugw(msg, keysAndValues...)
	})
}

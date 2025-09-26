package log

import (
	"fmt"
	"testing"
)

func TestWarnw_1(t *testing.T) {
	t.Run("test case 1:", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		msg := "test message"
		keysAndValues := []any{
			"key1", "value1",
			"key2", "value2",
		}
		Warnw(msg, keysAndValues...)
	})
}

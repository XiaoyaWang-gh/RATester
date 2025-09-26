package log

import (
	"fmt"
	"testing"
)

func TestPanicw_1(t *testing.T) {
	t.Parallel()
	t.Run("Panicw", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()
		msg := "msg"
		keysAndValues := []any{
			"key1", "value1",
			"key2", "value2",
		}
		Panicw(msg, keysAndValues...)
	})
}

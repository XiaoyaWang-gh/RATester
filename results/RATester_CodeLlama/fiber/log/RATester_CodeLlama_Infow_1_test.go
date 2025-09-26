package log

import (
	"fmt"
	"testing"
)

func TestInfow_1(t *testing.T) {
	t.Run("test Infow", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		msg := "test Infow"
		keysAndValues := []any{
			"key1", "value1",
			"key2", "value2",
		}
		Infow(msg, keysAndValues...)
	})
}

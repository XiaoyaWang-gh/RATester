package log

import (
	"fmt"
	"testing"
)

func TestDebugf_1(t *testing.T) {
	t.Run("TestDebugf", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		format := "test"
		v := []any{"test"}
		Debugf(format, v...)
	})
}

package log

import (
	"fmt"
	"testing"
)

func TestFatalw_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	msg := "msg"
	keysAndValues := []any{"key", "value"}
	Fatalw(msg, keysAndValues...)
}

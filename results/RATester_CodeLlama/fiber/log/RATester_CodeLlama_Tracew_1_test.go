package log

import (
	"fmt"
	"testing"
)

func TestTracew_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	msg := "test"
	keysAndValues := []any{"key", "value"}
	Tracew(msg, keysAndValues...)
}

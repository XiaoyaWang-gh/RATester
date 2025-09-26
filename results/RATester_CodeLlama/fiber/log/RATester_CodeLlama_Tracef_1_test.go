package log

import (
	"fmt"
	"testing"
)

func TestTracef_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	format := "test"
	v := []any{1, 2, 3}
	Tracef(format, v...)
}

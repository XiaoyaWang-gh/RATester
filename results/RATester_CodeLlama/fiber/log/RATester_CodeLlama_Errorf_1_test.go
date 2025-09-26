package log

import (
	"fmt"
	"testing"
)

func TestErrorf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	format := "test"
	v := []any{1, 2, 3}
	Errorf(format, v...)
}

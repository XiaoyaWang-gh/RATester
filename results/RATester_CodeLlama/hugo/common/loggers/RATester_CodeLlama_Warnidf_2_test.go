package loggers

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestWarnidf_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	id := "test"
	format := "test %s"
	v := []any{"test"}
	l.Warnidf(id, format, v...)
	assert.Equal(t, "test test", l.errors.String())
}

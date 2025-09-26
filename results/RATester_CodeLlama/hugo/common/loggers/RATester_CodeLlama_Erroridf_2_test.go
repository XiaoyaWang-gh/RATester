package loggers

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestErroridf_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &logAdapter{}
	id := "test"
	format := "test %s"
	v := []any{"test"}
	l.Erroridf(id, format, v...)
	assert.Equal(t, "test", id)
	assert.Equal(t, "test test", format)
	assert.Equal(t, []any{"test"}, v)
}

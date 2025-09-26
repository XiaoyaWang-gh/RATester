package loggers

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestOut_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	l := &logAdapter{}
	// when
	out := l.Out()
	// then
	assert.Equal(t, l.out, out)
}

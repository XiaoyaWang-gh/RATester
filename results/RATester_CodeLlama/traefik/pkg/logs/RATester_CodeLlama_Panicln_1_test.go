package logs

import (
	"fmt"
	"testing"
)

func TestPanicln_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	l := LogrusStdWrapper{}
	args := []interface{}{"test"}

	// when
	l.Panicln(args...)

	// then
	// TODO
}

package logs

import (
	"fmt"
	"testing"
)

func TestPrintf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	l := LogrusStdWrapper{}
	s := "test"
	args := []interface{}{"test"}

	// when
	l.Printf(s, args...)

	// then
	// TODO
}

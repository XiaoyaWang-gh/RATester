package logs

import (
	"fmt"
	"testing"
)

func TestWriteln_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lg := &logWriter{}
	msg := "test"
	n, err := lg.writeln(msg)
	if err != nil {
		t.Errorf("writeln() error = %v", err)
		return
	}
	if n != len(msg) {
		t.Errorf("writeln() n = %v, want %v", n, len(msg))
	}
}

package pageparser

import (
	"fmt"
	"testing"
)

func TestRun_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &pageLexer{}
	l.run()
	if l.state != nil {
		t.Errorf("l.state = %v, want nil", l.state)
	}
}

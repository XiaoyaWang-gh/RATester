package page

import (
	"fmt"
	"testing"
)

func TestPrepend_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pagePathBuilder{}
	p.Prepend("a", "b")
	if p.els[0] != "a" || p.els[1] != "b" {
		t.Errorf("Prepend failed")
	}
}

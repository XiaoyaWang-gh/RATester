package templates

import (
	"fmt"
	"testing"
)

func TestDefer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	if _, err := ns.Defer(); err == nil {
		t.Errorf("Defer does not take any arguments")
	}
}

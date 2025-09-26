package constraints

import (
	"fmt"
	"testing"
)

func TestAndLabelFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := func(labels map[string]string) bool {
		return true
	}
	b := func(labels map[string]string) bool {
		return true
	}
	c := andLabelFunc(a, b)
	if !c(map[string]string{}) {
		t.Errorf("expected true")
	}
}

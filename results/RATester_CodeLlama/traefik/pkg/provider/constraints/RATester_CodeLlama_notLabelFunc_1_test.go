package constraints

import (
	"fmt"
	"testing"
)

func TestNotLabelFunc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := func(labels map[string]string) bool {
		return true
	}
	b := notLabelFunc(a)
	if b(map[string]string{"a": "b"}) {
		t.Error("notLabelFunc failed")
	}
}

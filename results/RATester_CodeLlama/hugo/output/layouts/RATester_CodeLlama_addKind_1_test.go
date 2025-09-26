package layouts

import (
	"fmt"
	"testing"
)

func TestAddKind_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &layoutBuilder{}
	l.addKind()
	if len(l.layoutVariations) != 1 {
		t.Errorf("addKind() failed")
	}
}

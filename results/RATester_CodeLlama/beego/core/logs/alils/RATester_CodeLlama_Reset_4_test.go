package alils

import (
	"fmt"
	"testing"
)

func TestReset_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Log{}
	m.Reset()
	if m.Time != nil {
		t.Errorf("m.Time should be nil")
	}
	if m.Contents != nil {
		t.Errorf("m.Contents should be nil")
	}
}

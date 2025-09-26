package mirror

import (
	"fmt"
	"testing"
)

func TestDone_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := contextStopPropagation{}
	if c.Done() == nil {
		t.Errorf("Done() should not be nil")
	}
}

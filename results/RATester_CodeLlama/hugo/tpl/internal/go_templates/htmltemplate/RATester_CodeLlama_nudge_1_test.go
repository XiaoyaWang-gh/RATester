package template

import (
	"fmt"
	"testing"
)

func TestNudge_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := context{state: stateTag}
	c = nudge(c)
	if c.state != stateAttrName {
		t.Errorf("nudge(%v) = %v, want %v", c, c.state, stateAttrName)
	}
}

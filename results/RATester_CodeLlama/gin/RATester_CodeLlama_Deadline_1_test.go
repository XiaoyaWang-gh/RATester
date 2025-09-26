package gin

import (
	"fmt"
	"testing"
	"time"
)

func TestDeadline_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{}
	deadline, ok := c.Deadline()
	if deadline != (time.Time{}) {
		t.Errorf("Deadline() = %v, want %v", deadline, time.Time{})
	}
	if ok != false {
		t.Errorf("Deadline() = %v, want %v", ok, false)
	}
}

package fiber

import (
	"fmt"
	"testing"
	"time"
)

func TestSetDeadline_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var c testConn
	if err := c.SetDeadline(time.Time{}); err != nil {
		t.Errorf("SetDeadline() = %v, want nil", err)
	}
}

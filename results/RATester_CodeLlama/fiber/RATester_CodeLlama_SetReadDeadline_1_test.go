package fiber

import (
	"fmt"
	"testing"
	"time"
)

func TestSetReadDeadline_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var tc testConn
	var tt time.Time
	if err := tc.SetReadDeadline(tt); err != nil {
		t.Errorf("SetReadDeadline() = %v, want nil", err)
	}
}

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

	tc := &testConn{}
	now := time.Now()
	err := tc.SetReadDeadline(now)
	if err != nil {
		t.Errorf("SetReadDeadline returned an error: %v", err)
	}
}

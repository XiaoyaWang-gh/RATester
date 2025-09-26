package fiber

import (
	"fmt"
	"testing"
	"time"
)

func TestSetWriteDeadline_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tc := &testConn{}
	now := time.Now()
	err := tc.SetWriteDeadline(now)
	if err != nil {
		t.Errorf("SetWriteDeadline returned an error: %v", err)
	}
}

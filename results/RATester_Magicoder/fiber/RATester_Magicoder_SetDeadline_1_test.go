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

	tc := &testConn{}
	now := time.Now()
	err := tc.SetDeadline(now)
	if err != nil {
		t.Errorf("SetDeadline returned an error: %v", err)
	}
}

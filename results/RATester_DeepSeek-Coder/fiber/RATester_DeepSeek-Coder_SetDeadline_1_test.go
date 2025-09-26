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

	t.Parallel()

	conn := &testConn{}

	testTime := time.Now()
	err := conn.SetDeadline(testTime)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

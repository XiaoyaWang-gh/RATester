package fiber

import (
	"fmt"
	"testing"
)

func TestClose_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &testConn{}
	if err := conn.Close(); err != nil {
		t.Errorf("Close() = %v, want nil", err)
	}
}

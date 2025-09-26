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
	err := conn.Close()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

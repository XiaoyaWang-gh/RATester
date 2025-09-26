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

	// arrange
	var conn testConn
	var deadline time.Time
	// act
	err := conn.SetWriteDeadline(deadline)
	// assert
	if err != nil {
		t.Errorf("SetWriteDeadline() error = %v", err)
		return
	}
}

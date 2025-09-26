package rate

import (
	"fmt"
	"testing"
)

func TestClose_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &rateConn{}
	err := conn.Close()
	if err != nil {
		t.Errorf("Close() = %v, want nil", err)
	}
}

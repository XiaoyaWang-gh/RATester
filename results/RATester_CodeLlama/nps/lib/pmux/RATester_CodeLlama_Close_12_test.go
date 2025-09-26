package pmux

import (
	"fmt"
	"testing"
)

func TestClose_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	var (
		conn = &PortConn{}
		err  error
	)
	if err = conn.Close(); err != nil {
		t.Error(err)
	}
}

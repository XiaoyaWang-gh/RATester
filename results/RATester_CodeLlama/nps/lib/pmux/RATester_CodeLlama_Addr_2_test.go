package pmux

import (
	"fmt"
	"testing"
)

func TestAddr_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pListener := &PortListener{}
	addr := pListener.Addr()
	if addr == nil {
		t.Errorf("Addr() failed")
	}
}

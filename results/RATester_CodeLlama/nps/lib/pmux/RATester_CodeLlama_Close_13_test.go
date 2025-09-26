package pmux

import (
	"fmt"
	"testing"
)

func TestClose_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pListener := &PortListener{
		connCh: make(chan *PortConn, 10),
	}
	err := pListener.Close()
	if err != nil {
		t.Errorf("pListener.Close() = %v, want nil", err)
	}
	if pListener.isClose != true {
		t.Errorf("pListener.isClose = %v, want true", pListener.isClose)
	}
}

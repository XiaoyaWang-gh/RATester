package client

import (
	"fmt"
	"testing"
)

func TestDone_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &Control{}
	doneCh := ctl.Done()
	if doneCh == nil {
		t.Fatal("ctl.Done() should not be nil")
	}
}

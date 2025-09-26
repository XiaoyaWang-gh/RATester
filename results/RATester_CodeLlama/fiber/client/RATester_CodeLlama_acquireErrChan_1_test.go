package client

import (
	"fmt"
	"testing"
)

func TestAcquireErrChan_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ch := acquireErrChan()
	if ch == nil {
		t.Error("acquireErrChan() returned nil")
	}
}

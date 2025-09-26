package client

import (
	"fmt"
	"testing"
)

func TestReleaseErrChan_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ch := make(chan error)
	releaseErrChan(ch)

	if errChanPool.Get() != ch {
		t.Errorf("Expected channel to be returned from pool")
	}
}

package client

import (
	"fmt"
	"testing"
)

func TestreleaseErrChan_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ch := make(chan error)
	releaseErrChan(ch)

	// Test if the channel is in the pool
	_, ok := errChanPool.Get().(chan error)
	if !ok {
		t.Error("Expected a chan error, but got something else")
	}
}

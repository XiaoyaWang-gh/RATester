package acme

import (
	"fmt"
	"testing"
)

func TestCleanChan_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	challengeTLSALPN := &ChallengeTLSALPN{
		chans: make(map[string]chan struct{}),
	}

	key := "testKey"
	challengeTLSALPN.chans[key] = make(chan struct{})

	challengeTLSALPN.cleanChan(key)

	if _, ok := challengeTLSALPN.chans[key]; ok {
		t.Errorf("Expected channel to be closed and deleted, but it still exists")
	}
}

package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRun_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &nps{
		exit: make(chan struct{}),
	}

	go func() {
		time.Sleep(time.Second)
		close(n.exit)
	}()

	err := n.run()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

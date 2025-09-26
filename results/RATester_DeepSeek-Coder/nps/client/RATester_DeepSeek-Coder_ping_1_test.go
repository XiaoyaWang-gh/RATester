package client

import (
	"fmt"
	"testing"
	"time"
)

func TestPing_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &TRPClient{
		ticker: time.NewTicker(time.Second * 5),
	}

	go client.ping()

	time.Sleep(time.Second * 6)

	if client.ticker == nil {
		t.Errorf("Expected ticker to be not nil")
	}

	client.Close()

	time.Sleep(time.Second * 1)

	if client.ticker != nil {
		t.Errorf("Expected ticker to be nil after calling Close")
	}
}

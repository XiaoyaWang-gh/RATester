package client

import (
	"fmt"
	"testing"
	"time"
)

func TestGracefulClose_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{
		gracefulShutdownDuration: 10 * time.Second,
	}

	svr.GracefulClose(10 * time.Second)

	if svr.gracefulShutdownDuration != 10*time.Second {
		t.Errorf("Expected gracefulShutdownDuration to be 10 seconds, but got %v", svr.gracefulShutdownDuration)
	}
}

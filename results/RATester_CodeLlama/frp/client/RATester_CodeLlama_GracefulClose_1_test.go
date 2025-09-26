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

	svr := &Service{}
	svr.GracefulClose(time.Second)
	if svr.gracefulShutdownDuration != time.Second {
		t.Fatal("GracefulClose failed")
	}
}

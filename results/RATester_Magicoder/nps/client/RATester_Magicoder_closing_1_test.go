package client

import (
	"fmt"
	"testing"
)

func Testclosing_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &TRPClient{
		svrAddr: "test_addr",
	}

	client.closing()

	if !CloseClient {
		t.Error("CloseClient should be true")
	}

	if NowStatus != 0 {
		t.Error("NowStatus should be 0")
	}

	if client.tunnel != nil {
		t.Error("tunnel should be nil")
	}

	if client.signal != nil {
		t.Error("signal should be nil")
	}

	if client.ticker != nil {
		t.Error("ticker should be nil")
	}
}

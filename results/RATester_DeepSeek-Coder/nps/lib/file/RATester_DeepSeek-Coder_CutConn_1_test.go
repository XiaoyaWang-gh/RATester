package file

import (
	"fmt"
	"testing"
)

func TestCutConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		NowConn: 0,
	}

	client.CutConn()

	if client.NowConn != 1 {
		t.Errorf("Expected NowConn to be 1, got %d", client.NowConn)
	}
}

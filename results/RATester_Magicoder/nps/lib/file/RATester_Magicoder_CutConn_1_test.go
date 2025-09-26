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
		NowConn: 10,
	}

	client.CutConn()

	if client.NowConn != 11 {
		t.Errorf("Expected NowConn to be 11, but got %d", client.NowConn)
	}
}

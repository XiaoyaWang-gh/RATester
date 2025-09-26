package file

import (
	"fmt"
	"testing"
)

func TestAddConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		NowConn: 10,
	}

	client.AddConn()

	if client.NowConn != 9 {
		t.Errorf("Expected NowConn to be 9, but got %d", client.NowConn)
	}
}

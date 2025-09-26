package file

import (
	"fmt"
	"testing"
)

func TestGetConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		MaxConn: 10,
		NowConn: 5,
	}

	result := client.GetConn()

	if result != true {
		t.Errorf("Expected true, but got %v", result)
	}

	if client.NowConn != 6 {
		t.Errorf("Expected NowConn to be 6, but got %v", client.NowConn)
	}
}

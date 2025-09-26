package file

import (
	"fmt"
	"testing"
)

func TestGetTunnelNum_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		Id: 1,
	}

	GetDb().JsonDb.Tasks.Store(1, &Tunnel{
		Id:     1,
		Client: client,
	})

	GetDb().JsonDb.Tasks.Store(2, &Tunnel{
		Id:     2,
		Client: client,
	})

	num := client.GetTunnelNum()
	if num != 2 {
		t.Errorf("Expected 2, but got %d", num)
	}
}

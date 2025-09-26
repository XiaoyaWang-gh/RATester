package conn

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestGetTaskInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{
		Conn: nil,
		Rb:   []byte{},
	}

	expectedTunnel := &file.Tunnel{
		Id:      int(file.GetDb().JsonDb.GetTaskId()),
		NoStore: true,
		Flow:    new(file.Flow),
	}

	actualTunnel, err := conn.GetTaskInfo()

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if actualTunnel.Id != expectedTunnel.Id {
		t.Errorf("Expected Id: %v, but got: %v", expectedTunnel.Id, actualTunnel.Id)
	}

	if actualTunnel.NoStore != expectedTunnel.NoStore {
		t.Errorf("Expected NoStore: %v, but got: %v", expectedTunnel.NoStore, actualTunnel.NoStore)
	}

	if actualTunnel.Flow != expectedTunnel.Flow {
		t.Errorf("Expected Flow: %v, but got: %v", expectedTunnel.Flow, actualTunnel.Flow)
	}
}

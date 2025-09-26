package conn

import (
	"fmt"
	"testing"
)

func TestGetConfigInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{
		Conn: nil,
		Rb:   []byte{},
	}

	client, err := conn.GetConfigInfo()
	if err != nil {
		t.Errorf("GetConfigInfo() error = %v", err)
		return
	}

	if client.NoStore != true {
		t.Errorf("GetConfigInfo() NoStore = %v, want %v", client.NoStore, true)
	}

	if client.Status != true {
		t.Errorf("GetConfigInfo() Status = %v, want %v", client.Status, true)
	}

	if client.Flow == nil {
		t.Errorf("GetConfigInfo() Flow = %v, want not nil", client.Flow)
	}

	if client.NoDisplay != false {
		t.Errorf("GetConfigInfo() NoDisplay = %v, want %v", client.NoDisplay, false)
	}
}

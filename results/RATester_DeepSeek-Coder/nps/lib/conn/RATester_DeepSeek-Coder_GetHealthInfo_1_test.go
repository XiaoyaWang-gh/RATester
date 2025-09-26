package conn

import (
	"fmt"
	"testing"
)

func TestGetHealthInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{
		Conn: nil,
		Rb:   []byte("info|true"),
	}

	info, status, err := conn.GetHealthInfo()
	if err != nil {
		t.Errorf("GetHealthInfo() error = %v, wantErr %v", err, nil)
		return
	}
	if info != "info" {
		t.Errorf("GetHealthInfo() info = %v, want %v", info, "info")
	}
	if status != true {
		t.Errorf("GetHealthInfo() status = %v, want %v", status, true)
	}
}

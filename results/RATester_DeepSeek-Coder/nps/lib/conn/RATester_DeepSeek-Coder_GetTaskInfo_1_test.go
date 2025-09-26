package conn

import (
	"fmt"
	"testing"
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

	_, err := conn.GetTaskInfo()
	if err != nil {
		t.Errorf("GetTaskInfo() error = %v, wantErr %v", err, nil)
		return
	}
}

package conn

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestGetHostInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{}
	host, err := conn.GetHostInfo()
	if err != nil {
		t.Errorf("GetHostInfo failed, err: %v", err)
		return
	}
	if host == nil {
		t.Errorf("GetHostInfo failed, host is nil")
		return
	}
	if host.Id != int(file.GetDb().JsonDb.GetHostId()) {
		t.Errorf("GetHostInfo failed, host id is not correct")
		return
	}
	if host.Flow == nil {
		t.Errorf("GetHostInfo failed, host flow is nil")
		return
	}
	if host.NoStore != true {
		t.Errorf("GetHostInfo failed, host NoStore is not true")
		return
	}
}

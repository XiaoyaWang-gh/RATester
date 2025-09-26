package conn

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/common"
)

func TestWriteConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conn := &Conn{
		Conn: nil,
		Rb:   []byte(common.WORK_CONFIG),
	}

	n, err := conn.WriteConfig()
	if err != nil {
		t.Errorf("WriteConfig() error = %v, wantErr %v", err, nil)
		return
	}

	if n != len(common.WORK_CONFIG) {
		t.Errorf("WriteConfig() = %v, want %v", n, len(common.WORK_CONFIG))
	}
}

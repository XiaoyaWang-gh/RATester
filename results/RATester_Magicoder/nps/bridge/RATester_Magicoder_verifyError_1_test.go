package bridge

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/conn"
)

func TestverifyError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	bridge := &Bridge{}
	conn := &conn.Conn{}
	bridge.verifyError(conn)
}

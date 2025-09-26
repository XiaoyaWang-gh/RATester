package visitor

import (
	"fmt"
	"testing"

	netpkg "github.com/fatedier/frp/pkg/util/net"
)

func TestCloseListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	vm := &Manager{
		listeners: map[string]*listenerBundle{
			"test": {
				l:          &netpkg.InternalListener{},
				sk:         "test_sk",
				allowUsers: []string{"test_user"},
			},
		},
	}

	vm.CloseListener("test")

	if _, ok := vm.listeners["test"]; ok {
		t.Error("Expected listener to be deleted")
	}
}

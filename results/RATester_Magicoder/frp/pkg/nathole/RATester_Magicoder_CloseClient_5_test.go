package nathole

import (
	"fmt"
	"testing"
)

func TestCloseClient_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		clientCfgs: map[string]*ClientCfg{
			"test": {
				name: "test",
				sk:   "test_sk",
			},
		},
	}

	ctrl.CloseClient("test")

	if _, ok := ctrl.clientCfgs["test"]; ok {
		t.Error("Failed to close client")
	}
}

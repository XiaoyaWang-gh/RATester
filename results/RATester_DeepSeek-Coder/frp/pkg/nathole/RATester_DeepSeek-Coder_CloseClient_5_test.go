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

	t.Parallel()

	ctrl := &Controller{
		clientCfgs: map[string]*ClientCfg{
			"test": {
				name:       "test",
				sk:         "test_sk",
				allowUsers: []string{"user1", "user2"},
				sidCh:      make(chan string),
			},
		},
	}

	ctrl.CloseClient("test")

	if _, ok := ctrl.clientCfgs["test"]; ok {
		t.Errorf("Expected clientCfgs to be deleted, but it's still there")
	}
}

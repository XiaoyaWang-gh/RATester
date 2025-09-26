package nathole

import (
	"fmt"
	"testing"
)

func TestListenClient_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		clientCfgs: make(map[string]*ClientCfg),
	}

	name := "test"
	sk := "test_sk"
	allowUsers := []string{"user1", "user2"}

	_, err := ctrl.ListenClient(name, sk, allowUsers)
	if err != nil {
		t.Errorf("ListenClient failed: %v", err)
	}

	_, err = ctrl.ListenClient(name, sk, allowUsers)
	if err == nil {
		t.Errorf("ListenClient should have failed")
	}
}

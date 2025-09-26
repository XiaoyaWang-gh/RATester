package nathole

import (
	"fmt"
	"testing"
)

func TestListenClient_1(t *testing.T) {
	ctrl := &Controller{
		clientCfgs: make(map[string]*ClientCfg),
	}

	name := "test"
	sk := "test_sk"
	allowUsers := []string{"user1", "user2"}

	t.Run("normal case", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		sidCh, err := ctrl.ListenClient(name, sk, allowUsers)
		if err != nil {
			t.Fatalf("ListenClient failed: %v", err)
		}
		if sidCh == nil {
			t.Fatalf("ListenClient failed: sidCh is nil")
		}
		if _, ok := ctrl.clientCfgs[name]; !ok {
			t.Fatalf("ListenClient failed: clientCfgs[%s] not exist", name)
		}
	})

	t.Run("duplicate name case", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ctrl.ListenClient(name, "another_sk", allowUsers)
		if err == nil {
			t.Fatalf("ListenClient should fail")
		}
		if err.Error() != "proxy [test] is repeated" {
			t.Fatalf("ListenClient failed: err is not expected: %v", err)
		}
	})
}

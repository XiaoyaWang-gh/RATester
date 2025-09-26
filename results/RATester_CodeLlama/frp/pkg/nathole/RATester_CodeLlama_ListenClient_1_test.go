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

	c := &Controller{
		clientCfgs: make(map[string]*ClientCfg),
	}
	name := "test"
	sk := "test"
	allowUsers := []string{"test"}
	_, err := c.ListenClient(name, sk, allowUsers)
	if err != nil {
		t.Error(err)
	}
}

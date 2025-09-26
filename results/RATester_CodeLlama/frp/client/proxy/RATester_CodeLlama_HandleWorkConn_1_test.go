package proxy

import (
	"fmt"
	"testing"
)

func TestHandleWorkConn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pm := &Manager{
		proxies: make(map[string]*Wrapper),
	}
	pm.proxies["test"] = &Wrapper{}
	pm.HandleWorkConn("test", nil, nil)
}

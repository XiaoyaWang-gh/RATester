package proxy

import (
	"fmt"
	"testing"
)

func TestStartProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pm := &Manager{
		proxies: make(map[string]*Wrapper),
	}
	pm.proxies["test"] = &Wrapper{}
	err := pm.StartProxy("test", "127.0.0.1:8080", "test")
	if err != nil {
		t.Error(err)
	}
}

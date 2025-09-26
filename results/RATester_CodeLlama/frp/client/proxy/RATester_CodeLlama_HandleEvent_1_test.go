package proxy

import (
	"fmt"
	"testing"
)

func TestHandleEvent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pm := &Manager{
		proxies: make(map[string]*Wrapper),
	}
	pm.HandleEvent(nil)
}

package proxy

import (
	"fmt"
	"testing"
)

func TestNewTCPMuxProxy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	baseProxy := &BaseProxy{}
	proxy := NewTCPMuxProxy(baseProxy)
	if proxy == nil {
		t.Error("NewTCPMuxProxy failed")
	}
}

package proxy

import (
	"fmt"
	"testing"
)

func TestStop_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pw := &Wrapper{
		closeCh: make(chan struct{}),
	}
	pw.Stop()
	if pw.Phase != ProxyPhaseClosed {
		t.Errorf("pw.Phase = %s, want %s", pw.Phase, ProxyPhaseClosed)
	}
}

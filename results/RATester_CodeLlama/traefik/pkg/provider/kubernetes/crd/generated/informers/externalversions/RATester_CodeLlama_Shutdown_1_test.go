package externalversions

import (
	"fmt"
	"testing"
)

func TestShutdown_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := NewSharedInformerFactory(nil, 0)
	f.Shutdown()
	f.Shutdown()
}

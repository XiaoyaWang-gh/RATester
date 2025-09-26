package process

import (
	"fmt"
	"testing"
)

func TestStop_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Process{}
	p.SetBeforeStopHandler(func() {})
	p.Start()
	p.Stop()
	if p.stopped != true {
		t.Errorf("Process.Stop() failed")
	}
}

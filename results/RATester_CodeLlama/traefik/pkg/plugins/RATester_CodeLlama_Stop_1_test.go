package plugins

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

	p := _PP{}
	p.WStop = func() error {
		return nil
	}
	err := p.Stop()
	if err != nil {
		t.Errorf("Stop() error = %v", err)
		return
	}
}

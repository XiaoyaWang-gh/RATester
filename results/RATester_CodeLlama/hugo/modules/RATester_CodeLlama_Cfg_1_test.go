package modules

import (
	"fmt"
	"testing"
)

func TestCfg_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &moduleAdapter{}
	if got := m.Cfg(); got != nil {
		t.Errorf("Cfg() = %v, want nil", got)
	}
}

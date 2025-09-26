package modules

import (
	"fmt"
	"testing"
)

func TestGetMain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	modules := goModules{
		&goModule{Main: true},
		&goModule{Main: false},
		&goModule{Main: false},
	}

	mainModule := modules.GetMain()

	if mainModule == nil {
		t.Errorf("Expected main module, got nil")
	}

	if !mainModule.Main {
		t.Errorf("Expected main module, got non-main module")
	}
}

package identity

import (
	"fmt"
	"testing"
)

func TestGetDependencyManager_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := DependencyManagerProviderFunc(func() Manager {
		return nil
	})
	if got := d.GetDependencyManager(); got == nil {
		t.Errorf("DependencyManagerProviderFunc.GetDependencyManager() = %v, want %v", got, nil)
	}
}

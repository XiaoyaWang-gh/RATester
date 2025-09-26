package proxy

import (
	"fmt"
	"testing"
)

func TestNewManager_3(t *testing.T) {
	t.Run("NewManager", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		manager := NewManager()
		if manager == nil {
			t.Errorf("NewManager() = %v, want not nil", manager)
		}
		if manager.pxys == nil {
			t.Errorf("NewManager().pxys = %v, want not nil", manager.pxys)
		}
	})
}

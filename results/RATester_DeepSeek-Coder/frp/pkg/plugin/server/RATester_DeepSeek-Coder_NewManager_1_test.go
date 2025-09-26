package plugin

import (
	"fmt"
	"testing"
)

func TestNewManager_1(t *testing.T) {
	t.Run("TestNewManager", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		m := NewManager()
		if m == nil {
			t.Errorf("NewManager() = %v, want not nil", m)
		}
		if len(m.loginPlugins) != 0 {
			t.Errorf("NewManager().loginPlugins = %v, want empty", m.loginPlugins)
		}
	})
}

package session

import (
	"fmt"
	"testing"
)

func TestCfgEnableSidInURLQuery_1(t *testing.T) {
	t.Run("TestCfgEnableSidInURLQuery", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		config := &ManagerConfig{}
		CfgEnableSidInURLQuery(true)(config)
		if !config.EnableSidInURLQuery {
			t.Errorf("Expected EnableSidInURLQuery to be true, got false")
		}
		CfgEnableSidInURLQuery(false)(config)
		if config.EnableSidInURLQuery {
			t.Errorf("Expected EnableSidInURLQuery to be false, got true")
		}
	})
}

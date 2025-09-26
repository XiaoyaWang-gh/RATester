package session

import (
	"fmt"
	"testing"
)

func TestCfgGcLifeTime_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var config ManagerConfig
	config.Opts(CfgGcLifeTime(100))
	if config.Gclifetime != 100 {
		t.Errorf("config.Gclifetime = %d, want %d", config.Gclifetime, 100)
	}
}

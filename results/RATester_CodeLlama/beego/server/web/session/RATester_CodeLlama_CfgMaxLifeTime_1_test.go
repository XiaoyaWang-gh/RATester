package session

import (
	"fmt"
	"testing"
)

func TestCfgMaxLifeTime_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var lifeTime int64 = 100
	opt := CfgMaxLifeTime(lifeTime)
	config := &ManagerConfig{}
	opt(config)
	if config.Maxlifetime != lifeTime {
		t.Errorf("config.Maxlifetime = %d, want %d", config.Maxlifetime, lifeTime)
	}
}

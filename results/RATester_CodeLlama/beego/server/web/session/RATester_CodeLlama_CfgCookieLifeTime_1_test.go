package session

import (
	"fmt"
	"testing"
)

func TestCfgCookieLifeTime_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var lifeTime int
	var config ManagerConfig
	config.Opts(CfgCookieLifeTime(lifeTime))
	if config.CookieLifeTime != lifeTime {
		t.Errorf("config.CookieLifeTime = %d, want %d", config.CookieLifeTime, lifeTime)
	}
}

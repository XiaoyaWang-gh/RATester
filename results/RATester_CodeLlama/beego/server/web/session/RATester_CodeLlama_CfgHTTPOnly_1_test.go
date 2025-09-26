package session

import (
	"fmt"
	"testing"
)

func TestCfgHTTPOnly_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var config ManagerConfig
	config.Opts(CfgHTTPOnly(true))
	if config.DisableHTTPOnly != true {
		t.Errorf("CfgHTTPOnly failed")
	}
}

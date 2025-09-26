package session

import (
	"fmt"
	"testing"
)

func TestCfgSessionIdInHTTPHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	enable := true
	// act
	config := &ManagerConfig{}
	CfgSessionIdInHTTPHeader(enable)(config)
	// assert
	if config.EnableSidInHTTPHeader != enable {
		t.Errorf("Expected %v, but got %v", enable, config.EnableSidInHTTPHeader)
	}
}

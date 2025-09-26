package session

import (
	"fmt"
	"testing"
)

func TestCfgSessionIdPrefix_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	prefix := "test_prefix"
	opt := CfgSessionIdPrefix(prefix)
	config := &ManagerConfig{}
	opt(config)
	if config.SessionIDPrefix != prefix {
		t.Errorf("Expected SessionIDPrefix to be %s, got %s", prefix, config.SessionIDPrefix)
	}
}

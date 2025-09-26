package session

import (
	"fmt"
	"testing"
)

func TestCfgSetSessionNameInHTTPHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "test"
	opt := CfgSetSessionNameInHTTPHeader(name)
	config := &ManagerConfig{}
	opt(config)
	if config.SessionNameInHTTPHeader != name {
		t.Errorf("config.SessionNameInHTTPHeader = %v, want %v", config.SessionNameInHTTPHeader, name)
	}
}

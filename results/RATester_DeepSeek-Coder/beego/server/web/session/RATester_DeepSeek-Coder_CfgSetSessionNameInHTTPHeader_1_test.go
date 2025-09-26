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

	testName := "testSessionName"
	opt := CfgSetSessionNameInHTTPHeader(testName)
	config := &ManagerConfig{}
	opt(config)
	if config.SessionNameInHTTPHeader != testName {
		t.Errorf("Expected SessionNameInHTTPHeader to be %s, got %s", testName, config.SessionNameInHTTPHeader)
	}
}

package httplib

import (
	"fmt"
	"testing"
)

func TestSetDefaultSetting_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	testSetting := BeegoHTTPSettings{
		UserAgent: "TestUserAgent",
		// other fields omitted for brevity
	}

	SetDefaultSetting(testSetting)

	settingMutex.Lock()
	defer settingMutex.Unlock()

	if defaultSetting.UserAgent != testSetting.UserAgent {
		t.Errorf("Expected UserAgent to be %s, got %s", testSetting.UserAgent, defaultSetting.UserAgent)
	}

	// additional assertions for other fields omitted for brevity
}

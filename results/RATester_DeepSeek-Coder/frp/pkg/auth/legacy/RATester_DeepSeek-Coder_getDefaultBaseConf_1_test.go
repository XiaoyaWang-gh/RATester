package legacy

import (
	"fmt"
	"testing"
)

func TestGetDefaultBaseConf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := BaseConfig{
		AuthenticationMethod:     "token",
		AuthenticateHeartBeats:   false,
		AuthenticateNewWorkConns: false,
	}

	result := getDefaultBaseConf()

	if result.AuthenticationMethod != expected.AuthenticationMethod {
		t.Errorf("Expected AuthenticationMethod to be %v, got %v", expected.AuthenticationMethod, result.AuthenticationMethod)
	}

	if result.AuthenticateHeartBeats != expected.AuthenticateHeartBeats {
		t.Errorf("Expected AuthenticateHeartBeats to be %v, got %v", expected.AuthenticateHeartBeats, result.AuthenticateHeartBeats)
	}

	if result.AuthenticateNewWorkConns != expected.AuthenticateNewWorkConns {
		t.Errorf("Expected AuthenticateNewWorkConns to be %v, got %v", expected.AuthenticateNewWorkConns, result.AuthenticateNewWorkConns)
	}
}

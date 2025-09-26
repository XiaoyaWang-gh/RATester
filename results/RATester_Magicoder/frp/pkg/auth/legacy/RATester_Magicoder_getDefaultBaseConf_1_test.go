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
	if result != expected {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
}

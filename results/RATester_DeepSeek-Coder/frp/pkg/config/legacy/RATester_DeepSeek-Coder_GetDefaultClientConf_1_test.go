package legacy

import (
	"fmt"
	"testing"
)

func TestGetDefaultClientConf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conf := GetDefaultClientConf()

	if conf.ServerAddr != "0.0.0.0" {
		t.Errorf("Expected ServerAddr to be '0.0.0.0', got %s", conf.ServerAddr)
	}
}

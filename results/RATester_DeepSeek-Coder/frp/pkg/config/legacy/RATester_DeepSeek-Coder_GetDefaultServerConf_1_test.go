package legacy

import (
	"fmt"
	"testing"
)

func TestGetDefaultServerConf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conf := GetDefaultServerConf()

	if conf.DashboardAddr != "0.0.0.0" {
		t.Errorf("Expected DashboardAddr to be '0.0.0.0', got %s", conf.DashboardAddr)
	}
}

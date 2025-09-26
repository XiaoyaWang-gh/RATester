package session

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCfgSameSite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var sameSite http.SameSite
	var config *ManagerConfig
	var opt ManagerConfigOpt
	var expected ManagerConfigOpt

	// TEST CASE 1:
	sameSite = http.SameSiteLaxMode
	config = &ManagerConfig{}
	opt = CfgSameSite(sameSite)
	expected = func(config *ManagerConfig) {
		config.CookieSameSite = sameSite
	}
	config.Opts(opt)
	if !reflect.DeepEqual(config, expected) {
		t.Errorf("config is not equal to expected")
	}

	// TEST CASE 2:
	sameSite = http.SameSiteStrictMode
	config = &ManagerConfig{}
	opt = CfgSameSite(sameSite)
	expected = func(config *ManagerConfig) {
		config.CookieSameSite = sameSite
	}
	config.Opts(opt)
	if !reflect.DeepEqual(config, expected) {
		t.Errorf("config is not equal to expected")
	}

	// TEST CASE 3:
	sameSite = http.SameSiteNoneMode
	config = &ManagerConfig{}
	opt = CfgSameSite(sameSite)
	expected = func(config *ManagerConfig) {
		config.CookieSameSite = sameSite
	}
	config.Opts(opt)
	if !reflect.DeepEqual(config, expected) {
		t.Errorf("config is not equal to expected")
	}
}

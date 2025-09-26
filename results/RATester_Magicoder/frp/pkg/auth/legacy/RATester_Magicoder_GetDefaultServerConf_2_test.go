package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetDefaultServerConf_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := ServerConfig{
		BaseConfig:       getDefaultBaseConf(),
		OidcServerConfig: getDefaultOidcServerConf(),
		TokenConfig:      getDefaultTokenConf(),
	}

	actual := GetDefaultServerConf()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

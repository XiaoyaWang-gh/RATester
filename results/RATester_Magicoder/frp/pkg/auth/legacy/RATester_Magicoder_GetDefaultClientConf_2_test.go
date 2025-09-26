package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetDefaultClientConf_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := ClientConfig{
		BaseConfig:       getDefaultBaseConf(),
		OidcClientConfig: getDefaultOidcClientConf(),
		TokenConfig:      getDefaultTokenConf(),
	}

	actual := GetDefaultClientConf()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

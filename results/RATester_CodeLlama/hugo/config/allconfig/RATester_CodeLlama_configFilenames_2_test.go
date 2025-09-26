package allconfig

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfigFilenames_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := ConfigSourceDescriptor{
		Filename: "config.toml",
	}
	if got := d.configFilenames(); !reflect.DeepEqual(got, []string{"config.toml"}) {
		t.Errorf("configFilenames() = %v, want %v", got, []string{"config.toml"})
	}
}

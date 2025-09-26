package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetBaseConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cfg := &BaseVisitorConf{
		ProxyName:      "test",
		ProxyType:      "test",
		UseEncryption:  true,
		UseCompression: true,
		Role:           "test",
		Sk:             "test",
		ServerUser:     "test",
		ServerName:     "test",
		BindAddr:       "test",
		BindPort:       1,
	}
	if got := cfg.GetBaseConfig(); !reflect.DeepEqual(got, cfg) {
		t.Errorf("GetBaseConfig() = %v, want %v", got, cfg)
	}
}

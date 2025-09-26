package legacy

import (
	"fmt"
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
		BindPort:       8080,
	}

	result := cfg.GetBaseConfig()

	if result != cfg {
		t.Errorf("Expected %v, got %v", cfg, result)
	}
}

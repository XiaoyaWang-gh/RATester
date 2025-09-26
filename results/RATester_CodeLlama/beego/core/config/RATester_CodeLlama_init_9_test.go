package config

import (
	"fmt"
	"testing"
)

func TestInit_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	Register("ini", &IniConfig{})

	err := InitGlobalInstance("ini", "conf/app.conf")
	if err != nil {
		t.Errorf("init global config instance failed. If you do not use this, just ignore it. %v", err)
	}
}

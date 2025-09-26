package config

import (
	"fmt"
	"testing"
)

func TestInitGlobalInstance_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var err error
	name := "test"
	cfg := "test.cfg"
	err = InitGlobalInstance(name, cfg)
	if err != nil {
		t.Errorf("InitGlobalInstance failed, err: %v", err)
	}
}

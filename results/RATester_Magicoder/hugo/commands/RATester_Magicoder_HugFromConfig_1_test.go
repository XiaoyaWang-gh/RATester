package commands

import (
	"fmt"
	"testing"
)

func TestHugFromConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &rootCommand{}
	conf := &commonConfig{}

	h, err := r.HugFromConfig(conf)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if h == nil {
		t.Error("Expected a HugoSites instance, but got nil")
	}
}

package hugolib

import (
	"fmt"
	"testing"
)

func TestinitSites_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HugoSites{}
	config := &BuildCfg{}

	err := h.initSites(config)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

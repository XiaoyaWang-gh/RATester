package common

import (
	"fmt"
	"testing"
)

func TestGetConfigPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := GetConfigPath()
	if path != "conf/npc.conf" {
		t.Errorf("GetConfigPath() = %v, want %v", path, "conf/npc.conf")
	}
}

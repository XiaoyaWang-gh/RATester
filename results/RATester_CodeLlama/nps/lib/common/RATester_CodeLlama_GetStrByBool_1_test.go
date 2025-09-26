package common

import (
	"fmt"
	"testing"
)

func TestGetStrByBool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if GetStrByBool(true) != "1" {
		t.Error("GetStrByBool(true) != \"1\"")
	}
	if GetStrByBool(false) != "0" {
		t.Error("GetStrByBool(false) != \"0\"")
	}
}

package common

import (
	"fmt"
	"testing"
)

func TestGetExternalIp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	externalIp = ""
	ip := GetExternalIp()
	if ip == "" {
		t.Error("GetExternalIp() failed")
	}
}

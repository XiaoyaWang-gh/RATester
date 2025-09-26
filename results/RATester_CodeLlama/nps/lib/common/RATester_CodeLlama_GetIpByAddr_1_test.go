package common

import (
	"fmt"
	"testing"
)

func TestGetIpByAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	addr := "192.168.1.1:8080"
	ip := GetIpByAddr(addr)
	if ip != "192.168.1.1" {
		t.Errorf("GetIpByAddr(%s) = %s, want %s", addr, ip, "192.168.1.1")
	}
}

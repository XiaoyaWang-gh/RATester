package common

import (
	"fmt"
	"testing"
)

func TestGetIntranetIp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	err, ip := GetIntranetIp()
	if err != nil {
		t.Errorf("GetIntranetIp() error = %v, wantErr %v", err, nil)
		return
	}
	if ip == "" {
		t.Errorf("GetIntranetIp() ip = %v, want not empty", ip)
	}
}

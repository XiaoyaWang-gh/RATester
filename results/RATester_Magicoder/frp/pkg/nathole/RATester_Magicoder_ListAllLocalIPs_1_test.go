package nathole

import (
	"fmt"
	"testing"
)

func TestListAllLocalIPs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ips, err := ListAllLocalIPs()
	if err != nil {
		t.Errorf("ListAllLocalIPs() failed, error: %v", err)
	}

	for _, ip := range ips {
		if ip == nil {
			t.Errorf("ListAllLocalIPs() returned nil IP")
		}
		if ip.IsLoopback() {
			t.Errorf("ListAllLocalIPs() returned loopback IP: %v", ip)
		}
		if ip.IsPrivate() {
			t.Errorf("ListAllLocalIPs() returned private IP: %v", ip)
		}
		if ip.IsUnspecified() {
			t.Errorf("ListAllLocalIPs() returned unspecified IP: %v", ip)
		}
	}
}

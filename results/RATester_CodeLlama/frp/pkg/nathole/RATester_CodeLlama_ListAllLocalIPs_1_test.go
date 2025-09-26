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
		t.Fatal(err)
	}
	if len(ips) == 0 {
		t.Fatal("no local IPs found")
	}
	for _, ip := range ips {
		if !ip.IsPrivate() {
			t.Errorf("IP %s is not private", ip)
		}
	}
}

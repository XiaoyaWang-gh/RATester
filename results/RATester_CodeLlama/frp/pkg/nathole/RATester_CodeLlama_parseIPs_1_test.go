package nathole

import (
	"fmt"
	"testing"
)

func TestParseIPs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	addrs := []string{"127.0.0.1:8080", "127.0.0.1:8081", "127.0.0.1:8082"}
	ips := parseIPs(addrs)
	if len(ips) != 3 {
		t.Errorf("parseIPs() = %v, want %v", len(ips), 3)
	}
	if ips[0] != "127.0.0.1" {
		t.Errorf("parseIPs() = %v, want %v", ips[0], "127.0.0.1")
	}
	if ips[1] != "127.0.0.1" {
		t.Errorf("parseIPs() = %v, want %v", ips[1], "127.0.0.1")
	}
	if ips[2] != "127.0.0.1" {
		t.Errorf("parseIPs() = %v, want %v", ips[2], "127.0.0.1")
	}
}

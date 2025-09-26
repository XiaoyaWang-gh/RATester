package gin

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestprepareTrustedCIDRs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := &Engine{
		trustedProxies: []string{"192.168.0.1", "10.0.0.0/8"},
	}

	cidr, err := engine.prepareTrustedCIDRs()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(cidr) != 2 {
		t.Errorf("Expected 2 CIDRs, got %d", len(cidr))
	}

	_, ipnet, _ := net.ParseCIDR("192.168.0.1/32")
	if !reflect.DeepEqual(cidr[0], ipnet) {
		t.Errorf("Expected CIDR 192.168.0.1/32, got %v", cidr[0])
	}

	_, ipnet, _ = net.ParseCIDR("10.0.0.0/8")
	if !reflect.DeepEqual(cidr[1], ipnet) {
		t.Errorf("Expected CIDR 10.0.0.0/8, got %v", cidr[1])
	}
}

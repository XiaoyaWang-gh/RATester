package metrics

import (
	"fmt"
	"testing"
)

func TestHasServerURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dynamicConfig{
		services: map[string]map[string]bool{
			"service1": {
				"server1": true,
				"server2": true,
			},
			"service2": {
				"server3": true,
				"server4": true,
			},
		},
	}
	if !d.hasServerURL("service1", "server1") {
		t.Errorf("expected true, got false")
	}
	if d.hasServerURL("service1", "server3") {
		t.Errorf("expected false, got true")
	}
	if d.hasServerURL("service3", "server1") {
		t.Errorf("expected false, got true")
	}
}

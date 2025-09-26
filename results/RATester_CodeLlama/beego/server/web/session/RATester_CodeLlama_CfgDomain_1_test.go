package session

import (
	"fmt"
	"testing"
)

func TestCfgDomain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	domain := "example.com"
	opt := CfgDomain(domain)
	config := &ManagerConfig{}
	opt(config)
	if config.Domain != domain {
		t.Errorf("config.Domain = %s, want %s", config.Domain, domain)
	}
}

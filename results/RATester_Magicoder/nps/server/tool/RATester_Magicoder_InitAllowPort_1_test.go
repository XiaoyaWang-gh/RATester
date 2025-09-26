package tool

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego"
)

func TestInitAllowPort_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	beego.AppConfig.Set("allow_ports", "80,443,8080")
	InitAllowPort()
	if len(ports) != 3 {
		t.Errorf("Expected 3 ports, got %d", len(ports))
	}
	if ports[0] != 80 || ports[1] != 443 || ports[2] != 8080 {
		t.Errorf("Expected ports [80, 443, 8080], got %v", ports)
	}
}

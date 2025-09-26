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

	beego.AppConfig.Set("allow_ports", "8080,8081,8082")
	InitAllowPort()
	if len(ports) != 3 {
		t.Errorf("Expected 3 ports, got %d", len(ports))
	}
	if ports[0] != 8080 || ports[1] != 8081 || ports[2] != 8082 {
		t.Errorf("Ports are not as expected, got %v", ports)
	}
}

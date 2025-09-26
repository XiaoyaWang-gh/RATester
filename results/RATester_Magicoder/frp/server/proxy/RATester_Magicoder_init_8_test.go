package proxy

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestInit_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	RegisterProxyFactory(reflect.TypeOf(&v1.XTCPProxyConfig{}), NewXTCPProxy)

	// Add assertions here to validate the behavior of the function under test
}

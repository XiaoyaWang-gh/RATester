package proxy

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestRegisterProxyFactory_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	RegisterProxyFactory(reflect.TypeOf(&BaseProxy{}), func(*BaseProxy, v1.ProxyConfigurer) Proxy {
		return nil
	})
}

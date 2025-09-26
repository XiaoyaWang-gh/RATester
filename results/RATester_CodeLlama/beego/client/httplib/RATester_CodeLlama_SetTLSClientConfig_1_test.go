package httplib

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestSetTLSClientConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := new(BeegoHTTPRequest)
	config := new(tls.Config)
	b.SetTLSClientConfig(config)
	if b.setting.TLSClientConfig != config {
		t.Error("TestSetTLSClientConfig failed")
	}
}

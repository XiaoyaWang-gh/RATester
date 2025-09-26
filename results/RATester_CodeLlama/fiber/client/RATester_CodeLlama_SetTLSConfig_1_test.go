package client

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestSetTLSConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Client{}
	config := &tls.Config{}
	c.SetTLSConfig(config)
	if c.fasthttp.TLSConfig != config {
		t.Errorf("SetTLSConfig failed")
	}
}

package fiber

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func Testprefork_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}
	addr := ":3000"
	tlsConfig := &tls.Config{}
	cfg := ListenConfig{}

	err := app.prefork(addr, tlsConfig, cfg)
	if err != nil {
		t.Errorf("prefork() error = %v", err)
	}
}

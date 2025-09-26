package fiber

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestClientHelloInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		tlsHandler: &TLSHandler{
			clientHelloInfo: &tls.ClientHelloInfo{},
		},
	}

	ctx := &DefaultCtx{
		app: app,
	}

	result := ctx.ClientHelloInfo()

	if result != app.tlsHandler.clientHelloInfo {
		t.Errorf("Expected %v, got %v", app.tlsHandler.clientHelloInfo, result)
	}
}

package fiber

import (
	"crypto/tls"
	"fmt"
	"testing"
)

func TestGetClientInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	handler := &TLSHandler{}
	info := &tls.ClientHelloInfo{}
	cert, err := handler.GetClientInfo(info)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if cert != nil {
		t.Errorf("Expected nil, got %v", cert)
	}
}

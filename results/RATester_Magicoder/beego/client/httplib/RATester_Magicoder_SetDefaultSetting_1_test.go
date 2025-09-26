package httplib

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestSetDefaultSetting_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	setting := BeegoHTTPSettings{
		UserAgent:        "TestAgent",
		ConnectTimeout:   10 * time.Second,
		ReadWriteTimeout: 15 * time.Second,
		TLSClientConfig:  &tls.Config{},
		Proxy:            http.ProxyFromEnvironment,
		Transport:        &http.Transport{},
		CheckRedirect:    nil,
		EnableCookie:     true,
		Gzip:             true,
		Retries:          3,
		RetryDelay:       5 * time.Second,
		FilterChains:     []FilterChain{},
		EscapeHTML:       true,
	}
	SetDefaultSetting(setting)

	// Add assertions here to verify the setting was set correctly
}

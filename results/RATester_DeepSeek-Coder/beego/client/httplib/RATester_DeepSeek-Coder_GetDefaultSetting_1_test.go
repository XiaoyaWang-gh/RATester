package httplib

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestGetDefaultSetting_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	expected := BeegoHTTPSettings{
		UserAgent:        "Go-http-client/1.1",
		ConnectTimeout:   30 * time.Second,
		ReadWriteTimeout: 30 * time.Second,
		TLSClientConfig:  &tls.Config{InsecureSkipVerify: true},
		Proxy:            http.ProxyFromEnvironment,
		Transport:        &http.Transport{},
		CheckRedirect:    nil,
		EnableCookie:     true,
		Gzip:             true,
		Retries:          3,
		RetryDelay:       500 * time.Millisecond,
		FilterChains:     []FilterChain{},
		EscapeHTML:       true,
	}

	result := GetDefaultSetting()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

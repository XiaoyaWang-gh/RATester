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
		UserAgent:        "Beego",
		ConnectTimeout:   10 * time.Second,
		ReadWriteTimeout: 10 * time.Second,
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

	actual := GetDefaultSetting()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}

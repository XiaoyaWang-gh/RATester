package httplib

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDefaultSetting_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	setting := GetDefaultSetting()
	if setting.UserAgent != "beego" {
		t.Errorf("UserAgent should be beego")
	}
	if setting.ConnectTimeout != 30*time.Second {
		t.Errorf("ConnectTimeout should be 30s")
	}
	if setting.ReadWriteTimeout != 80*time.Second {
		t.Errorf("ReadWriteTimeout should be 80s")
	}
	if setting.TLSClientConfig != nil {
		t.Errorf("TLSClientConfig should be nil")
	}
	if setting.Proxy != nil {
		t.Errorf("Proxy should be nil")
	}
	if setting.Transport != nil {
		t.Errorf("Transport should be nil")
	}
	if setting.CheckRedirect != nil {
		t.Errorf("CheckRedirect should be nil")
	}
	if setting.EnableCookie != true {
		t.Errorf("EnableCookie should be true")
	}
	if setting.Gzip != true {
		t.Errorf("Gzip should be true")
	}
	if setting.Retries != 0 {
		t.Errorf("Retries should be 0")
	}
	if setting.RetryDelay != 0 {
		t.Errorf("RetryDelay should be 0")
	}
	if setting.FilterChains != nil {
		t.Errorf("FilterChains should be nil")
	}
	if setting.EscapeHTML != true {
		t.Errorf("EscapeHTML should be true")
	}
}

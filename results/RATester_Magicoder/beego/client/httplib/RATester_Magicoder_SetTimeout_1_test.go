package httplib

import (
	"fmt"
	"testing"
	"time"
)

func TestSetTimeout_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{}
	connectTimeout := time.Duration(10)
	readWriteTimeout := time.Duration(20)
	b.SetTimeout(connectTimeout, readWriteTimeout)

	if b.setting.ConnectTimeout != connectTimeout {
		t.Errorf("Expected connect timeout to be %v, but got %v", connectTimeout, b.setting.ConnectTimeout)
	}

	if b.setting.ReadWriteTimeout != readWriteTimeout {
		t.Errorf("Expected read/write timeout to be %v, but got %v", readWriteTimeout, b.setting.ReadWriteTimeout)
	}
}

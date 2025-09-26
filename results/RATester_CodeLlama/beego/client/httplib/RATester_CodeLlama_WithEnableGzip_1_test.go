package httplib

import (
	"fmt"
	"testing"
)

func TestWithEnableGzip_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	enable := true
	clientOption := WithEnableGzip(enable)
	client := &Client{}
	clientOption(client)
	if client.Setting.Gzip != enable {
		t.Errorf("WithEnableGzip failed")
	}
}

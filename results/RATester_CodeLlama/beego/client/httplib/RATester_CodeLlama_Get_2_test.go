package httplib

import (
	"fmt"
	"testing"
)

func TestGet_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	url := "http://www.baidu.com"
	req := Get(url)
	if req.url != url {
		t.Errorf("Get url error, excepted: %s, got: %s", url, req.url)
	}
}

package web

import (
	"fmt"
	"net/url"
	"testing"
)

func TestStore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fd := &FlashData{
		Data: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	}

	c := &Controller{
		Data: make(map[interface{}]interface{}),
	}

	fd.Store(c)

	cookie, err := c.Ctx.Request.Cookie(BConfig.WebConfig.FlashName)
	if err != nil {
		t.Errorf("Failed to get cookie: %v", err)
	}

	value, err := url.QueryUnescape(cookie.Value)
	if err != nil {
		t.Errorf("Failed to unescape cookie value: %v", err)
	}

	expected := "\x00key1\x23\x23value1\x00\x00key2\x23\x23value2\x00"
	if value != expected {
		t.Errorf("Expected %q, got %q", expected, value)
	}
}

package context

import (
	"fmt"
	"testing"
)

func TestXSRFToken_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{
		_xsrfToken: "",
	}

	key := "testKey"
	expire := int64(3600)

	token := ctx.XSRFToken(key, expire)

	if token == "" {
		t.Error("XSRFToken should not be empty")
	}

	if ctx._xsrfToken != token {
		t.Error("XSRFToken should be set in context")
	}
}

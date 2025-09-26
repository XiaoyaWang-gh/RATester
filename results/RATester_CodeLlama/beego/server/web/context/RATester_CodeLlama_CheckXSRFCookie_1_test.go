package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCheckXSRFCookie_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &Context{}
	ctx.Request = &http.Request{}
	ctx.ResponseWriter = &Response{}
	ctx.Input = &BeegoInput{}
	ctx.Output = &BeegoOutput{}
	ctx._xsrfToken = "token"
	ctx.Request.Header = make(http.Header)
	ctx.Request.Header.Set("X-Xsrftoken", "token")
	ctx.Request.Header.Set("X-Csrftoken", "token")
	if !ctx.CheckXSRFCookie() {
		t.Error("CheckXSRFCookie failed")
	}
}

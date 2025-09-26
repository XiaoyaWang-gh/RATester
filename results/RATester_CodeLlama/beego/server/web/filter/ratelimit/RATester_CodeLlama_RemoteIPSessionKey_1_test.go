package ratelimit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
)

func TestRemoteIPSessionKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &context.Context{}
	ctx.Request = &http.Request{}
	ctx.Request.Header = make(http.Header)
	ctx.Request.Header.Set("X-Real-Ip", "192.168.1.1")
	ctx.Request.Header.Set("X-Forwarded-For", "192.168.1.2")
	ctx.Request.RemoteAddr = "192.168.1.3"
	assert.Equal(t, "192.168.1.1", RemoteIPSessionKey(ctx))
}

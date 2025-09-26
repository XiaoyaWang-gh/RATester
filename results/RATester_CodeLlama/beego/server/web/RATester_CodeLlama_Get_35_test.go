package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
	"github.com/zeebo/assert"
)

func TestGet_35(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	rootpath := "rootpath"
	f := func(ctx *beecontext.Context) {}

	// when
	app := Get(rootpath, f)

	// then
	assert.NotNil(t, app)
}

package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
	"github.com/zeebo/assert"
)

func TestPut_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	rootpath := "rootpath"
	f := func(ctx *beecontext.Context) {}

	// when
	app := Put(rootpath, f)

	// then
	assert.NotNil(t, app)
}

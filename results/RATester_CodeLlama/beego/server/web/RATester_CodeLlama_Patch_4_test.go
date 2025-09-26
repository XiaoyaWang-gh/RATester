package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/require"
)

func TestPatch_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: setup test
	rootpath := "TODO"
	f := func(ctx *beecontext.Context) {
		// TODO: setup test
	}
	app := &HttpServer{}
	got := app.Patch(rootpath, f)
	require.Equal(t, app, got)
}

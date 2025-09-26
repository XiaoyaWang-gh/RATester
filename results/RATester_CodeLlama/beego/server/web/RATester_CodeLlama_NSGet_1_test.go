package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestNSGet_1(t *testing.T) {
	t.Run("test NSGet", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		rootpath := "rootpath"
		f := func(ctx *beecontext.Context) {}
		ns := NSGet(rootpath, f)
		if ns == nil {
			t.Errorf("NSGet() = %v, want %v", ns, "not nil")
		}
	})
}

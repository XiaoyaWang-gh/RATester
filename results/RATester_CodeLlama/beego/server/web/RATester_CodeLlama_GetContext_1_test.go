package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestGetContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &ControllerRegister{}
	p.pool.New = func() interface{} {
		return &beecontext.Context{}
	}
	ctx := p.GetContext()
	if ctx == nil {
		t.Errorf("TestGetContext failed")
	}
}

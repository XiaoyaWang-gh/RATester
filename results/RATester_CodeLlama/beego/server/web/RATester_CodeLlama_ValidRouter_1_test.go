package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestValidRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &FilterRouter{}
	url := "test"
	ctx := &context.Context{}
	if f.ValidRouter(url, ctx) {
		t.Errorf("ValidRouter() = %v, want %v", true, false)
	}
}

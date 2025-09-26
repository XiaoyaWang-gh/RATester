package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
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

	ctrl := &Controller{
		Ctx: &context.Context{},
		Data: map[interface{}]interface{}{
			"flash": fd.Data,
		},
	}

	fd.Store(ctrl)

	expected := "\x00key1\x23\x23value1\x00\x00key2\x23\x23value2\x00"
	actual := ctrl.Ctx.GetCookie(BConfig.WebConfig.FlashName)

	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

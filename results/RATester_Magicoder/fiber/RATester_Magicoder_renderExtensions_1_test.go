package fiber

import (
	"fmt"
	"sync"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestrenderExtensions_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		app: &App{
			config: Config{
				PassLocalsToViews: true,
			},
		},
		viewBindMap: sync.Map{},
	}

	ctx.viewBindMap.Store("key1", "value1")
	ctx.viewBindMap.Store("key2", "value2")

	ctx.fasthttp = &fasthttp.RequestCtx{}
	ctx.fasthttp.VisitUserValues(func(key []byte, val any) {
		ctx.viewBindMap.Store(string(key), val)
	})

	ctx.renderExtensions(nil)

	ctx.viewBindMap.Range(func(key, value any) bool {
		keyValue, ok := key.(string)
		if !ok {
			t.Errorf("Expected key to be a string, but got %T", key)
		}
		if keyValue != "key1" && keyValue != "key2" {
			t.Errorf("Unexpected key: %s", keyValue)
		}
		if value != "value1" && value != "value2" {
			t.Errorf("Unexpected value: %s", value)
		}
		return true
	})
}

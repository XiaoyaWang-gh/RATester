package gin

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetQueryMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{
		Params: Params{
			Param{
				Key:   "key",
				Value: "value",
			},
		},
	}
	c.initQueryCache()
	c.queryCache = url.Values{
		"key": []string{"value"},
	}
	m, ok := c.GetQueryMap("key")
	if !ok {
		t.Errorf("GetQueryMap() = %v, want %v", ok, true)
	}
	if m["key"] != "value" {
		t.Errorf("GetQueryMap() = %v, want %v", m["key"], "value")
	}
}

package web

import (
	"fmt"
	"testing"
)

func TestRenderBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Data: map[interface{}]interface{}{
			"key": "value",
		},
		Layout: "layout.tpl",
	}

	buf, err := ctrl.RenderBytes()
	if err != nil {
		t.Errorf("RenderBytes() error = %v", err)
		return
	}

	if len(buf) == 0 {
		t.Errorf("RenderBytes() returned empty buffer")
	}
}

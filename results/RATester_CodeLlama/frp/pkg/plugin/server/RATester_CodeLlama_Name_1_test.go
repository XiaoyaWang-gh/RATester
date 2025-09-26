package plugin

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &httpPlugin{
		options: v1.HTTPPluginOptions{
			Name: "test",
		},
	}
	if p.Name() != "test" {
		t.Errorf("expect test, but got %s", p.Name())
	}
}

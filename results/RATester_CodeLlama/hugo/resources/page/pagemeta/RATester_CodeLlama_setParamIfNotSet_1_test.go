package pagemeta

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestSetParamIfNotSet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &FrontMatterDescriptor{
		PageConfig: &PageConfig{
			Params: make(maps.Params),
		},
	}
	setParamIfNotSet("key", "value", d)
	if d.PageConfig.Params["key"] != "value" {
		t.Errorf("Expected %q to be %q", d.PageConfig.Params["key"], "value")
	}
}

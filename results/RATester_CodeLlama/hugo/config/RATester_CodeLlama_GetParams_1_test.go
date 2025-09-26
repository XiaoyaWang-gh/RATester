package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &defaultConfigProvider{}
	c.Set("a", maps.Params{"b": "c"})
	if got := c.GetParams("a"); !reflect.DeepEqual(got, maps.Params{"b": "c"}) {
		t.Errorf("GetParams() = %v, want %v", got, maps.Params{"b": "c"})
	}
}

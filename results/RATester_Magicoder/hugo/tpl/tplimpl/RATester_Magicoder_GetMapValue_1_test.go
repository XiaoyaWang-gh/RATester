package tplimpl

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetMapValue_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	helper := &templateExecHelper{
		siteParams: reflect.ValueOf(maps.Params{"key": "value"}),
	}

	key := reflect.ValueOf("key")
	value, found := helper.GetMapValue(ctx, nil, helper.siteParams, key)

	if !found {
		t.Error("Expected value to be found")
	}

	if value.String() != "value" {
		t.Errorf("Expected value to be 'value', but got %s", value.String())
	}
}

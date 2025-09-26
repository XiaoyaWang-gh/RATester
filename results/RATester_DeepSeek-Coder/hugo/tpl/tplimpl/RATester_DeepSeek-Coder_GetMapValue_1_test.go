package tplimpl

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestGetMapValue_1(t *testing.T) {
	type testStruct struct {
		Name string
	}

	ctx := context.Background()
	helper := &templateExecHelper{}

	t.Run("TestGetMapValue_Found", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		params := make(maps.Params)
		params["name"] = "test"
		helper.siteParams = reflect.ValueOf(params)
		key := reflect.ValueOf("name")
		val, found := helper.GetMapValue(ctx, nil, helper.siteParams, key)
		if !found {
			t.Errorf("Expected value to be found")
		}
		if val.String() != "test" {
			t.Errorf("Expected 'test', got %s", val.String())
		}
	})

	t.Run("TestGetMapValue_NotFound", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		params := make(maps.Params)
		params["name"] = "test"
		helper.siteParams = reflect.ValueOf(params)
		key := reflect.ValueOf("non-existing-key")
		_, found := helper.GetMapValue(ctx, nil, helper.siteParams, key)
		if found {
			t.Errorf("Expected value not to be found")
		}
	})

	t.Run("TestGetMapValue_InvalidKey", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		params := make(maps.Params)
		params["name"] = "test"
		helper.siteParams = reflect.ValueOf(params)
		key := reflect.ValueOf(123)
		_, found := helper.GetMapValue(ctx, nil, helper.siteParams, key)
		if found {
			t.Errorf("Expected value not to be found")
		}
	})
}

package param

import (
	"fmt"
	"reflect"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestConvertParams_1(t *testing.T) {
	t.Run("TestConvertParams", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var methodParams []*MethodParam
		var methodType reflect.Type
		var ctx *beecontext.Context
		var result []reflect.Value
		if got := ConvertParams(methodParams, methodType, ctx); !reflect.DeepEqual(got, result) {
			t.Errorf("ConvertParams() = %v, want %v", got, result)
		}
	})
}

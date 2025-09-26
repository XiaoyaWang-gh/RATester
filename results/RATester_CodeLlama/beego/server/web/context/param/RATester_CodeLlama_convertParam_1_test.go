package param

import (
	"fmt"
	"reflect"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestConvertParam_1(t *testing.T) {
	type args struct {
		param     *MethodParam
		paramType reflect.Type
		ctx       *beecontext.Context
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertParam(tt.args.param, tt.args.paramType, tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

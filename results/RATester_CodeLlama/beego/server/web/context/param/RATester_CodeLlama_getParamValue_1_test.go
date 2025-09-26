package param

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestGetParamValue_1(t *testing.T) {
	type args struct {
		param *MethodParam
		ctx   *beecontext.Context
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := getParamValue(tt.args.param, tt.args.ctx); got != tt.want {
				t.Errorf("getParamValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

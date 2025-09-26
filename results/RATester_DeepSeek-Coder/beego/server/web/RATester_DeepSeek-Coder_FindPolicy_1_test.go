package web

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestFindPolicy_1(t *testing.T) {
	type args struct {
		cont *context.Context
	}
	tests := []struct {
		name string
		p    *ControllerRegister
		args args
		want []PolicyFunc
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

			if got := tt.p.FindPolicy(tt.args.cont); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ControllerRegister.FindPolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

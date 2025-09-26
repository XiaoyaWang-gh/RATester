package authz

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web"
	"github.com/casbin/casbin"
)

func TestNewAuthorizer_1(t *testing.T) {
	type args struct {
		e *casbin.Enforcer
	}
	tests := []struct {
		name string
		args args
		want web.FilterFunc
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

			if got := NewAuthorizer(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthorizer() = %v, want %v", got, tt.want)
			}
		})
	}
}

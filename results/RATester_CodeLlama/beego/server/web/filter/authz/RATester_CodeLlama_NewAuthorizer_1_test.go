package authz

import (
	"fmt"
	"testing"

	"github.com/casbin/casbin"
)

func TestNewAuthorizer_1(t *testing.T) {
	type args struct {
		e *casbin.Enforcer
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				e: &casbin.Enforcer{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			NewAuthorizer(tt.args.e)
		})
	}
}

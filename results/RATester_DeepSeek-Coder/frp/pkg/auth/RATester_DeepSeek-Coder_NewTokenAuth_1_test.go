package auth

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewTokenAuth_1(t *testing.T) {
	type args struct {
		additionalAuthScopes []v1.AuthScope
		token                string
	}
	tests := []struct {
		name string
		args args
		want *TokenAuthSetterVerifier
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

			if got := NewTokenAuth(tt.args.additionalAuthScopes, tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTokenAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

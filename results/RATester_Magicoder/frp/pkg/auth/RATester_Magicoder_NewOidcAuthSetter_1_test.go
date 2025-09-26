package auth

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestNewOidcAuthSetter_1(t *testing.T) {
	type args struct {
		additionalAuthScopes []v1.AuthScope
		cfg                  v1.AuthOIDCClientConfig
	}
	tests := []struct {
		name string
		args args
		want *OidcAuthProvider
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

			if got := NewOidcAuthSetter(tt.args.additionalAuthScopes, tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOidcAuthSetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

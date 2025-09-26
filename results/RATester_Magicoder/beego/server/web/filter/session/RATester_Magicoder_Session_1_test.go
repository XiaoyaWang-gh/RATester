package session

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/session"
)

func TestSession_1(t *testing.T) {
	type args struct {
		providerType session.ProviderType
		options      []session.ManagerConfigOpt
	}
	tests := []struct {
		name string
		args args
		want web.FilterChain
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

			if got := Session(tt.args.providerType, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session() = %v, want %v", got, tt.want)
			}
		})
	}
}

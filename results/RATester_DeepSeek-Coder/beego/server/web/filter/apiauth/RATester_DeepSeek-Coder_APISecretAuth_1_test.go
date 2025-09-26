package apiauth

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web"
)

func TestAPISecretAuth_1(t *testing.T) {
	type args struct {
		f       AppIDToAppSecret
		timeout int
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

			if got := APISecretAuth(tt.args.f, tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("APISecretAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

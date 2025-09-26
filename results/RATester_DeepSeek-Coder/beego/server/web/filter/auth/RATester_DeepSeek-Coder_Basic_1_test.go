package auth

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web"
)

func TestBasic_1(t *testing.T) {
	type args struct {
		username string
		password string
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

			if got := Basic(tt.args.username, tt.args.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Basic() = %v, want %v", got, tt.want)
			}
		})
	}
}

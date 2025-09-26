package cors

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web"
)

func TestAllow_1(t *testing.T) {
	type args struct {
		opts *Options
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

			if got := Allow(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Allow() = %v, want %v", got, tt.want)
			}
		})
	}
}

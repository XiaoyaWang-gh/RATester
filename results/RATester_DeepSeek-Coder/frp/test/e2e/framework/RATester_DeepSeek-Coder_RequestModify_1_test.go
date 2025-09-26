package framework

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fatedier/frp/test/e2e/pkg/request"
)

func TestRequestModify_1(t *testing.T) {
	type args struct {
		f func(r *request.Request)
	}
	tests := []struct {
		name string
		e    *RequestExpect
		args args
		want *RequestExpect
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

			if got := tt.e.RequestModify(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RequestExpect.RequestModify() = %v, want %v", got, tt.want)
			}
		})
	}
}

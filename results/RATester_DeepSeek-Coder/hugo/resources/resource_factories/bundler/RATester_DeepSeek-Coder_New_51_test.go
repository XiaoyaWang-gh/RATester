package bundler

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources"
)

func TestNew_51(t *testing.T) {
	type args struct {
		rs *resources.Spec
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "TestNew",
			args: args{
				rs: &resources.Spec{},
			},
			want: &Client{
				rs: &resources.Spec{},
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

			if got := New(tt.args.rs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

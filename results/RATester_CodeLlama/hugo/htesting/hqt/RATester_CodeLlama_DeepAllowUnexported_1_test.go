package hqt

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDeepAllowUnexported_1(t *testing.T) {
	type args struct {
		vs []any
	}
	tests := []struct {
		name string
		args args
		want cmp.Option
	}{
		{
			name: "test case 1",
			args: args{
				vs: []any{
					struct{}{},
				},
			},
			want: cmp.AllowUnexported(struct{}{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := DeepAllowUnexported(tt.args.vs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeepAllowUnexported() = %v, want %v", got, tt.want)
			}
		})
	}
}

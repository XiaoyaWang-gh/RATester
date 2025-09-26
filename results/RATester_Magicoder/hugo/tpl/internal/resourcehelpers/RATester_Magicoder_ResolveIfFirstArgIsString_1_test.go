package resourcehelpers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources"
)

func TestResolveIfFirstArgIsString_1(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name  string
		args  args
		want  resources.ResourceTransformer
		want1 string
		want2 bool
	}{
		{
			name: "Test case 1",
			args: args{
				args: []any{"test", resources.ResourceTransformer(nil)},
			},
			want:  resources.ResourceTransformer(nil),
			want1: "test",
			want2: true,
		},
		{
			name: "Test case 2",
			args: args{
				args: []any{123, resources.ResourceTransformer(nil)},
			},
			want:  nil,
			want1: "",
			want2: false,
		},
		{
			name: "Test case 3",
			args: args{
				args: []any{"test"},
			},
			want:  nil,
			want1: "",
			want2: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1, got2 := ResolveIfFirstArgIsString(tt.args.args)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveIfFirstArgIsString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ResolveIfFirstArgIsString() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ResolveIfFirstArgIsString() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

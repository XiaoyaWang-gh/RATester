package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIndirectToStringerOrError_1(t *testing.T) {
	type args struct {
		a any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "test1",
			args: args{
				a: "test",
			},
			want: "test",
		},
		{
			name: "test2",
			args: args{
				a: &struct {
					Name string
				}{
					Name: "test",
				},
			},
			want: &struct {
				Name string
			}{
				Name: "test",
			},
		},
		{
			name: "test3",
			args: args{
				a: &struct {
					Name string
				}{
					Name: "test",
				},
			},
			want: &struct {
				Name string
			}{
				Name: "test",
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

			if got := indirectToStringerOrError(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("indirectToStringerOrError() = %v, want %v", got, tt.want)
			}
		})
	}
}

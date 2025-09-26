package redactor

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsExported_1(t *testing.T) {
	type args struct {
		f reflect.StructField
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				f: reflect.StructField{
					Name:      "field name",
					PkgPath:   "package path",
					Type:      reflect.TypeOf(1),
					Tag:       reflect.StructTag("tag"),
					Offset:    1,
					Index:     []int{1},
					Anonymous: true,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isExported(tt.args.f); got != tt.want {
				t.Errorf("isExported() = %v, want %v", got, tt.want)
			}
		})
	}
}

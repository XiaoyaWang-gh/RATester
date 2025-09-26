package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestformTagName_1(t *testing.T) {
	type args struct {
		fieldT reflect.StructField
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "Test case 1",
			args: args{
				fieldT: reflect.StructField{
					Name: "Name",
					Tag:  reflect.StructTag(`form:"name"`),
				},
			},
			want:  "name",
			want1: true,
		},
		{
			name: "Test case 2",
			args: args{
				fieldT: reflect.StructField{
					Name: "Name",
					Tag:  reflect.StructTag(`form:"-"`),
				},
			},
			want:  "",
			want1: false,
		},
		{
			name: "Test case 3",
			args: args{
				fieldT: reflect.StructField{
					Name: "Name",
					Tag:  reflect.StructTag(`form:""`),
				},
			},
			want:  "Name",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := formTagName(tt.args.fieldT)
			if got != tt.want {
				t.Errorf("formTagName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("formTagName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

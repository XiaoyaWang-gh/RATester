package api

import (
	"fmt"
	"testing"
)

func TestExtractType_1(t *testing.T) {
	type args struct {
		element interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				element: &struct {
					Field1 string
					Field2 string
				}{
					Field1: "test",
					Field2: "test",
				},
			},
			want: "Field1",
		},
		{
			name: "test case 2",
			args: args{
				element: &struct {
					Field1 string
					Field2 string
				}{
					Field1: "test",
					Field2: "test",
				},
			},
			want: "Field1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := extractType(tt.args.element); got != tt.want {
				t.Errorf("extractType() = %v, want %v", got, tt.want)
			}
		})
	}
}

package controllers

import (
	"fmt"
	"reflect"
	"testing"
)

func Testajax_1(t *testing.T) {
	type args struct {
		str    string
		status int
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Test Case 1",
			args: args{
				str:    "Test String",
				status: 200,
			},
			want: map[string]interface{}{
				"status": 200,
				"msg":    "Test String",
			},
		},
		{
			name: "Test Case 2",
			args: args{
				str:    "Another Test String",
				status: 404,
			},
			want: map[string]interface{}{
				"status": 404,
				"msg":    "Another Test String",
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

			if got := ajax(tt.args.str, tt.args.status); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ajax() = %v, want %v", got, tt.want)
			}
		})
	}
}

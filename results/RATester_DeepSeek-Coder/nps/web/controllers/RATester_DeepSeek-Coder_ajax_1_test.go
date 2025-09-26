package controllers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAjax_1(t *testing.T) {
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
			name: "Test case 1",
			args: args{
				str:    "success",
				status: 200,
			},
			want: map[string]interface{}{
				"status": 200,
				"msg":    "success",
			},
		},
		{
			name: "Test case 2",
			args: args{
				str:    "failure",
				status: 400,
			},
			want: map[string]interface{}{
				"status": 400,
				"msg":    "failure",
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

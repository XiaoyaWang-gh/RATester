package utils

import (
	"fmt"
	"testing"
)

func TestGetDisplayString_1(t *testing.T) {
	type args struct {
		data []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				data: []interface{}{"field name"},
			},
			want: "field want",
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetDisplayString(tt.args.data...); got != tt.want {
				t.Errorf("GetDisplayString() = %v, want %v", got, tt.want)
			}
		})
	}
}

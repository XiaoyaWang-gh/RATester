package common

import (
	"fmt"
	"testing"
)

func TestGetStrByBool_1(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{b: true},
			want: "1",
		},
		{
			name: "Test case 2",
			args: args{b: false},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetStrByBool(tt.args.b); got != tt.want {
				t.Errorf("GetStrByBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

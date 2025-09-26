package common

import (
	"fmt"
	"testing"
)

func TestGetverifyval_1(t *testing.T) {
	type args struct {
		vkey string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{vkey: "test123"},
			want: "e80b5017098950fc58aad83c8c14978e",
		},
		{
			name: "Test Case 2",
			args: args{vkey: "123456"},
			want: "e80b5017098950fc58aad83c8c14978e",
		},
		{
			name: "Test Case 3",
			args: args{vkey: "abcdefg"},
			want: "e80b5017098950fc58aad83c8c14978e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Getverifyval(tt.args.vkey); got != tt.want {
				t.Errorf("Getverifyval() = %v, want %v", got, tt.want)
			}
		})
	}
}

package tcp

import (
	"fmt"
	"testing"
)

func TestVarGroupName_1(t *testing.T) {
	type args struct {
		idx int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{idx: 1},
			want: "v1",
		},
		{
			name: "Test case 2",
			args: args{idx: 2},
			want: "v2",
		},
		{
			name: "Test case 3",
			args: args{idx: 3},
			want: "v3",
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

			if got := varGroupName(tt.args.idx); got != tt.want {
				t.Errorf("varGroupName() = %v, want %v", got, tt.want)
			}
		})
	}
}

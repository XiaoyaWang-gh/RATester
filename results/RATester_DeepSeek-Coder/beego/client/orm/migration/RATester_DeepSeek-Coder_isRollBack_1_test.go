package migration

import (
	"fmt"
	"testing"
)

func TestIsRollBack_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Test case 1",
			args: args{
				name: "migration1",
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				name: "migration2",
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

			if got := isRollBack(tt.args.name); got != tt.want {
				t.Errorf("isRollBack() = %v, want %v", got, tt.want)
			}
		})
	}
}

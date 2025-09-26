package fiber

import (
	"fmt"
	"testing"
)

func TestfindNextCharsetPositionConstraint_1(t *testing.T) {
	type args struct {
		search  string
		charset []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				search:  "test",
				charset: []byte{'t'},
			},
			want: 0,
		},
		{
			name: "Test Case 2",
			args: args{
				search:  "test",
				charset: []byte{'e'},
			},
			want: 1,
		},
		{
			name: "Test Case 3",
			args: args{
				search:  "test",
				charset: []byte{'s'},
			},
			want: 2,
		},
		{
			name: "Test Case 4",
			args: args{
				search:  "test",
				charset: []byte{'x'},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := findNextCharsetPositionConstraint(tt.args.search, tt.args.charset); got != tt.want {
				t.Errorf("findNextCharsetPositionConstraint() = %v, want %v", got, tt.want)
			}
		})
	}
}

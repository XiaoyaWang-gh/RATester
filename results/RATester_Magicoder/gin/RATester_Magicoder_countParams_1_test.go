package gin

import (
	"fmt"
	"testing"
)

func TestcountParams_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "Test case 1",
			args: args{
				path: "/api/v1/users/:id",
			},
			want: 1,
		},
		{
			name: "Test case 2",
			args: args{
				path: "/api/v1/users/*",
			},
			want: 1,
		},
		{
			name: "Test case 3",
			args: args{
				path: "/api/v1/users/:id/*",
			},
			want: 2,
		},
		{
			name: "Test case 4",
			args: args{
				path: "/api/v1/users",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := countParams(tt.args.path); got != tt.want {
				t.Errorf("countParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

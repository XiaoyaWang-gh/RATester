package gin

import (
	"fmt"
	"testing"
)

func TestCountParams_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			name: "testcase1",
			args: args{
				path: "/test/test1",
			},
			want: 0,
		},
		{
			name: "testcase2",
			args: args{
				path: "/test/:test1",
			},
			want: 1,
		},
		{
			name: "testcase3",
			args: args{
				path: "/test/*test1",
			},
			want: 1,
		},
		{
			name: "testcase4",
			args: args{
				path: "/test/:test1/:test2",
			},
			want: 2,
		},
		{
			name: "testcase5",
			args: args{
				path: "/test/*test1/*test2",
			},
			want: 2,
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

package urlreplacers

import (
	"fmt"
	"testing"
)

func Testfind_1(t *testing.T) {
	type args struct {
		bs    []byte
		start int
	}
	tests := []struct {
		name string
		p    *prefix
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			p: &prefix{
				disabled: false,
				b:        []byte("test"),
				f:        nil,
				nextPos:  -1,
			},
			args: args{
				bs:    []byte("teststring"),
				start: 0,
			},
			want: true,
		},
		{
			name: "Test Case 2",
			p: &prefix{
				disabled: false,
				b:        []byte("test"),
				f:        nil,
				nextPos:  -1,
			},
			args: args{
				bs:    []byte("otherstring"),
				start: 0,
			},
			want: false,
		},
		{
			name: "Test Case 3",
			p: &prefix{
				disabled: true,
				b:        []byte("test"),
				f:        nil,
				nextPos:  -1,
			},
			args: args{
				bs:    []byte("teststring"),
				start: 0,
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

			if got := tt.p.find(tt.args.bs, tt.args.start); got != tt.want {
				t.Errorf("prefix.find() = %v, want %v", got, tt.want)
			}
		})
	}
}

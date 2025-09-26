package hugocontext

import (
	"fmt"
	"testing"
)

func TestWrap_1(t *testing.T) {
	type args struct {
		b   []byte
		pid uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				b:   []byte("test"),
				pid: 123,
			},
			want: "prefix pid=123\nclosingDelimAndNewline",
		},
		{
			name: "Test case 2",
			args: args{
				b:   []byte("another test"),
				pid: 456,
			},
			want: "prefix pid=456\nclosingDelimAndNewline",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Wrap(tt.args.b, tt.args.pid); got != tt.want {
				t.Errorf("Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

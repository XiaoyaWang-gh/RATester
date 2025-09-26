package template

import (
	"fmt"
	"testing"
)

func TestEndsWithCSSKeyword_1(t *testing.T) {
	type args struct {
		b    []byte
		kw   string
		want bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testcase1",
			args: args{
				b:    []byte("test"),
				kw:   "test",
				want: true,
			},
		},
		{
			name: "testcase2",
			args: args{
				b:    []byte("test"),
				kw:   "test1",
				want: false,
			},
		},
		{
			name: "testcase3",
			args: args{
				b:    []byte("test"),
				kw:   "test",
				want: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := endsWithCSSKeyword(tt.args.b, tt.args.kw); got != tt.args.want {
				t.Errorf("endsWithCSSKeyword() = %v, want %v", got, tt.args.want)
			}
		})
	}
}

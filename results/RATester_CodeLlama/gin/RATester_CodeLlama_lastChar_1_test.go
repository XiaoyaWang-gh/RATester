package gin

import (
	"fmt"
	"testing"
)

func TestLastChar_1(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "test case 1",
			args: args{
				str: "hello",
			},
			want: 'o',
		},
		{
			name: "test case 2",
			args: args{
				str: "world",
			},
			want: 'd',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := lastChar(tt.args.str); got != tt.want {
				t.Errorf("lastChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

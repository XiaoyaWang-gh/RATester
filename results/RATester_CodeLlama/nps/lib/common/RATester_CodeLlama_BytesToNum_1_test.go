package common

import (
	"fmt"
	"testing"
)

func TestBytesToNum_1(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{b: []byte{1, 2, 3}},
			want: 123,
		},
		{
			name: "test case 2",
			args: args{b: []byte{1, 2, 3, 4}},
			want: 1234,
		},
		{
			name: "test case 3",
			args: args{b: []byte{1, 2, 3, 4, 5}},
			want: 12345,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := BytesToNum(tt.args.b); got != tt.want {
				t.Errorf("BytesToNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

package utils

import (
	"fmt"
	"testing"
)

func TestInSlice_1(t *testing.T) {
	type args struct {
		v  string
		sl []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				v:  "test",
				sl: []string{"test1", "test2", "test"},
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				v:  "test3",
				sl: []string{"test1", "test2", "test"},
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				v:  "",
				sl: []string{"test1", "test2", "test"},
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

			if got := InSlice(tt.args.v, tt.args.sl); got != tt.want {
				t.Errorf("InSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

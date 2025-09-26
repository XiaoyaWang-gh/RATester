package task

import (
	"fmt"
	"testing"
)

func TestmustParseInt_1(t *testing.T) {
	type args struct {
		expr string
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "Test case 1",
			args: args{expr: "10"},
			want: 10,
		},
		{
			name: "Test case 2",
			args: args{expr: "0"},
			want: 0,
		},
		{
			name: "Test case 3",
			args: args{expr: "-1"},
			want: 0,
		},
		{
			name: "Test case 4",
			args: args{expr: "abc"},
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

			if got := mustParseInt(tt.args.expr); got != tt.want {
				t.Errorf("mustParseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

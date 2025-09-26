package metrics

import (
	"fmt"
	"testing"
)

func TesthowSimilar_1(t *testing.T) {
	type args struct {
		a any
		b any
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				a: 1,
				b: 1,
			},
			want: 100,
		},
		{
			name: "Test case 2",
			args: args{
				a: "hello",
				b: "hello",
			},
			want: 100,
		},
		{
			name: "Test case 3",
			args: args{
				a: 1,
				b: "1",
			},
			want: 0,
		},
		{
			name: "Test case 4",
			args: args{
				a: "hello",
				b: 1,
			},
			want: 0,
		},
		{
			name: "Test case 5",
			args: args{
				a: 1,
				b: 2,
			},
			want: 0,
		},
		{
			name: "Test case 6",
			args: args{
				a: "hello",
				b: "world",
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

			if got := howSimilar(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("howSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}

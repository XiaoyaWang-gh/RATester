package compare

import (
	"fmt"
	"testing"
)

func TestLe_2(t *testing.T) {
	n := &Namespace{}
	type args struct {
		first  any
		others []any
	}
	tests := []struct {
		name string
		n    *Namespace
		args args
		want bool
	}{
		{
			name: "1",
			n:    n,
			args: args{
				first:  1,
				others: []any{2},
			},
			want: true,
		},
		{
			name: "2",
			n:    n,
			args: args{
				first:  1,
				others: []any{1},
			},
			want: true,
		},
		{
			name: "3",
			n:    n,
			args: args{
				first:  1,
				others: []any{0},
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

			if got := tt.n.Le(tt.args.first, tt.args.others...); got != tt.want {
				t.Errorf("Le() = %v, want %v", got, tt.want)
			}
		})
	}
}

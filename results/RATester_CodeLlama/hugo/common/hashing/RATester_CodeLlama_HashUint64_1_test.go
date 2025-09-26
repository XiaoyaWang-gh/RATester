package hashing

import (
	"fmt"
	"testing"
)

func TestHashUint64_1(t *testing.T) {
	type args struct {
		vs []any
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "empty",
			args: args{
				vs: []any{},
			},
			want: 0,
		},
		{
			name: "one",
			args: args{
				vs: []any{
					"foo",
				},
			},
			want: 1208071513,
		},
		{
			name: "two",
			args: args{
				vs: []any{
					"foo",
					"bar",
				},
			},
			want: 1208071513,
		},
		{
			name: "three",
			args: args{
				vs: []any{
					"foo",
					"bar",
					"baz",
				},
			},
			want: 1208071513,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HashUint64(tt.args.vs...); got != tt.want {
				t.Errorf("HashUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

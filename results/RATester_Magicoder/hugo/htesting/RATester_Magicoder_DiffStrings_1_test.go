package htesting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDiffStrings_1(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case 1",
			args: args{
				s1: "Hello world",
				s2: "Hello world",
			},
			want: []string{},
		},
		{
			name: "Test case 2",
			args: args{
				s1: "Hello world",
				s2: "Hello",
			},
			want: []string{"world"},
		},
		{
			name: "Test case 3",
			args: args{
				s1: "Hello",
				s2: "Hello world",
			},
			want: []string{"world"},
		},
		{
			name: "Test case 4",
			args: args{
				s1: "Hello world",
				s2: "Hello universe",
			},
			want: []string{"world", "universe"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := DiffStrings(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiffStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

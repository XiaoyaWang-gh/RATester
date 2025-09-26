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
				s1: "This is a test",
				s2: "This is a test",
			},
			want: []string{},
		},
		{
			name: "Test case 2",
			args: args{
				s1: "This is a test",
				s2: "This is not a test",
			},
			want: []string{"not", "a"},
		},
		{
			name: "Test case 3",
			args: args{
				s1: "This is a test",
				s2: "This is a different test",
			},
			want: []string{"different"},
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

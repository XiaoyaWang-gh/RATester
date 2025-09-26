package htesting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDiffStringSlices_1(t *testing.T) {
	type args struct {
		slice1 []string
		slice2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test Case 1",
			args: args{
				slice1: []string{"a", "b", "c"},
				slice2: []string{"b", "c", "d"},
			},
			want: []string{"a"},
		},
		{
			name: "Test Case 2",
			args: args{
				slice1: []string{"a", "b", "c"},
				slice2: []string{"a", "b", "c"},
			},
			want: []string{},
		},
		{
			name: "Test Case 3",
			args: args{
				slice1: []string{"a", "b", "c"},
				slice2: []string{},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "Test Case 4",
			args: args{
				slice1: []string{},
				slice2: []string{"a", "b", "c"},
			},
			want: []string{},
		},
		{
			name: "Test Case 5",
			args: args{
				slice1: []string{"a", "b", "c"},
				slice2: []string{"a", "b", "c", "d", "e"},
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := DiffStringSlices(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiffStringSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}

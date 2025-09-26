package htesting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDiffStringSlices_1(t *testing.T) {
	type test struct {
		slice1 []string
		slice2 []string
		want   []string
	}

	tests := []test{
		{
			slice1: []string{"a", "b", "c"},
			slice2: []string{"b", "c", "d"},
			want:   []string{"a", "d"},
		},
		{
			slice1: []string{"apple", "banana", "cherry"},
			slice2: []string{"banana", "cherry", "date"},
			want:   []string{"apple", "date"},
		},
		{
			slice1: []string{"1", "2", "3"},
			slice2: []string{"1", "2", "3", "4"},
			want:   []string{"4"},
		},
		{
			slice1: []string{},
			slice2: []string{"1", "2", "3"},
			want:   []string{"1", "2", "3"},
		},
		{
			slice1: []string{"1", "2", "3"},
			slice2: []string{},
			want:   []string{"1", "2", "3"},
		},
		{
			slice1: []string{},
			slice2: []string{},
			want:   []string{},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("DiffStringSlices(%v, %v)", tt.slice1, tt.slice2), func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := DiffStringSlices(tt.slice1, tt.slice2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

package port

import (
	"fmt"
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/util/sets"
)

func TestNewAllocator_1(t *testing.T) {
	type args struct {
		from  int
		to    int
		mod   int
		index int
	}
	tests := []struct {
		name string
		args args
		want *Allocator
	}{
		{
			name: "Test case 1",
			args: args{
				from:  1,
				to:    10,
				mod:   2,
				index: 0,
			},
			want: &Allocator{
				reserved: sets.New[int](),
				used:     sets.New[int](),
			},
		},
		{
			name: "Test case 2",
			args: args{
				from:  1,
				to:    10,
				mod:   3,
				index: 1,
			},
			want: &Allocator{
				reserved: sets.New[int](),
				used:     sets.New[int](),
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewAllocator(tt.args.from, tt.args.to, tt.args.mod, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAllocator() = %v, want %v", got, tt.want)
			}
		})
	}
}

package protoexample

import (
	"fmt"
	reflect "reflect"
	"testing"
)

func TestGetReps_1(t *testing.T) {
	tests := []struct {
		name string
		x    *Test
		want []int64
	}{
		{
			name: "nil",
			x:    nil,
			want: nil,
		},
		{
			name: "empty",
			x:    &Test{},
			want: nil,
		},
		{
			name: "non-empty",
			x: &Test{
				Reps: []int64{1, 2, 3},
			},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.x.GetReps(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReps() = %v, want %v", got, tt.want)
			}
		})
	}
}

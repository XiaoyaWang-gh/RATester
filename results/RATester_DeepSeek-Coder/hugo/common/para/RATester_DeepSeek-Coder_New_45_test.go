package para

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew_45(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		numWorkers int
		want       *Workers
	}{
		{
			name:       "Test case 1",
			numWorkers: 5,
			want: &Workers{
				sem: make(chan struct{}, 5),
			},
		},
		{
			name:       "Test case 2",
			numWorkers: 10,
			want: &Workers{
				sem: make(chan struct{}, 10),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := New(tt.numWorkers)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

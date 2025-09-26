package migration

import (
	"fmt"
	"testing"
)

func TestLess_3(t *testing.T) {
	tests := []struct {
		name string
		d    dataSlice
		i    int
		j    int
		want bool
	}{
		{
			name: "Test case 1",
			d:    dataSlice{data{created: 1, name: "test1", m: nil}, data{created: 2, name: "test2", m: nil}},
			i:    0,
			j:    1,
			want: true,
		},
		{
			name: "Test case 2",
			d:    dataSlice{data{created: 2, name: "test1", m: nil}, data{created: 1, name: "test2", m: nil}},
			i:    0,
			j:    1,
			want: false,
		},
		{
			name: "Test case 3",
			d:    dataSlice{data{created: 1, name: "test1", m: nil}, data{created: 1, name: "test2", m: nil}},
			i:    0,
			j:    1,
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

			if got := tt.d.Less(tt.i, tt.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

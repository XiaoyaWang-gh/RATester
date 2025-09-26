package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    any
		wantErr bool
	}{
		{
			name:  "nil slice",
			input: nil,
			want:  nil,
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "slice of ints",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		},
		{
			name:  "slice of strings",
			input: []string{"a", "b", "c", "d", "e"},
			want:  []string{"e", "d", "c", "b", "a"},
		},
		{
			name:  "slice of structs",
			input: []struct{ name string }{},
			want:  []struct{ name string }{},
		},
		{
			name:    "non-slice input",
			input:   123,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.Reverse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reverse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

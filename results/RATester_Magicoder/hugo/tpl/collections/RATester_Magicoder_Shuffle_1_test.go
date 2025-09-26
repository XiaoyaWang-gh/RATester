package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShuffle_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    any
		wantErr bool
	}{
		{
			name:    "nil input",
			input:   nil,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty slice",
			input:   []int{},
			want:    []int{},
			wantErr: false,
		},
		{
			name:    "slice of ints",
			input:   []int{1, 2, 3, 4, 5},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "slice of strings",
			input:   []string{"a", "b", "c", "d", "e"},
			want:    []string{"a", "b", "c", "d", "e"},
			wantErr: false,
		},
		{
			name:    "non-slice input",
			input:   "not a slice",
			want:    nil,
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

			got, err := ns.Shuffle(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Shuffle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}

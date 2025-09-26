package maps

import (
	"fmt"
	"testing"
)

func TestAdd_8(t *testing.T) {
	type testCase struct {
		name      string
		key       string
		newAddend any
		wantErr   bool
	}

	tests := []testCase{
		{
			name:      "Add to existing slice",
			key:       "sliceKey",
			newAddend: []any{1, 2, 3},
			wantErr:   false,
		},
		{
			name:      "Add to existing non-slice",
			key:       "nonSliceKey",
			newAddend: 4,
			wantErr:   false,
		},
		{
			name:      "Add to non-existing key",
			key:       "nonExistingKey",
			newAddend: "newValue",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &Scratch{
				values: make(map[string]any),
			}
			_, err := s.Add(tt.key, tt.newAddend)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scratch.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

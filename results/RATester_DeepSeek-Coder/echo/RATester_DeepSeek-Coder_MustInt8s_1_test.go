package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMustInt8s_1(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		source   string
		expected []int8
		wantErr  bool
	}

	tests := []testCase{
		{
			name:     "Test case 1",
			source:   "1,2,3",
			expected: []int8{1, 2, 3},
			wantErr:  false,
		},
		{
			name:     "Test case 2",
			source:   "1,2,3,256",
			expected: []int8{1, 2, 3},
			wantErr:  true,
		},
		{
			name:     "Test case 3",
			source:   "",
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			b := &ValueBinder{}
			dest := &[]int8{}
			err := b.MustInt8s(tt.source, dest).BindError()
			if (err != nil) != tt.wantErr {
				t.Errorf("MustInt8s() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(*dest, tt.expected) {
				t.Errorf("MustInt8s() = %v, want %v", *dest, tt.expected)
			}
		})
	}
}

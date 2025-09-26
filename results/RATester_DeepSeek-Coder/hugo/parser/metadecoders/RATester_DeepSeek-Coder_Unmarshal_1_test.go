package metadecoders

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestUnmarshal_1(t *testing.T) {
	type testCase struct {
		name     string
		data     []byte
		f        Format
		expected any
		err      error
	}

	tests := []testCase{
		{
			name:     "empty data and CSV format",
			data:     []byte{},
			f:        CSV,
			expected: make([][]string, 0),
			err:      nil,
		},
		{
			name:     "non-empty data and CSV format",
			data:     []byte("a,b,c\n1,2,3"),
			f:        CSV,
			expected: [][]string{{"a", "b", "c"}, {"1", "2", "3"}},
			err:      nil,
		},
		{
			name:     "empty data and default format",
			data:     []byte{},
			f:        "",
			expected: make(map[string]any),
			err:      nil,
		},
		{
			name:     "non-empty data and default format",
			data:     []byte("a:1,b:2,c:3"),
			f:        "",
			expected: map[string]any{"a": "1", "b": "2", "c": "3"},
			err:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := Decoder{}
			actual, err := d.Unmarshal(tt.data, tt.f)
			if !reflect.DeepEqual(actual, tt.expected) || !errors.Is(err, tt.err) {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}

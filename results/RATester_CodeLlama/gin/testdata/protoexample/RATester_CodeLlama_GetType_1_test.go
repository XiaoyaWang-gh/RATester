package protoexample

import (
	"fmt"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestGetType_1(t *testing.T) {
	tests := []struct {
		name     string
		input    *Test
		expected int32
	}{
		{
			name:     "TestGetType_1",
			input:    &Test{Type: proto.Int32(1)},
			expected: 1,
		},
		{
			name:     "TestGetType_2",
			input:    &Test{},
			expected: 77,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := test.input.GetType(); got != test.expected {
				t.Errorf("GetType() = %v, want %v", got, test.expected)
			}
		})
	}
}

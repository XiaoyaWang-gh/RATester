package warpc

import (
	"fmt"
	"testing"
)

func TestGetID_1(t *testing.T) {
	type testStruct struct {
		Field string `json:"field"`
	}

	testCases := []struct {
		name     string
		message  Message[testStruct]
		expected uint32
	}{
		{
			name: "Test case 1",
			message: Message[testStruct]{
				Header: Header{
					ID: 123,
				},
				Data: testStruct{
					Field: "test",
				},
			},
			expected: 123,
		},
		{
			name: "Test case 2",
			message: Message[testStruct]{
				Header: Header{
					ID: 456,
				},
				Data: testStruct{
					Field: "test",
				},
			},
			expected: 456,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.message.GetID()
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

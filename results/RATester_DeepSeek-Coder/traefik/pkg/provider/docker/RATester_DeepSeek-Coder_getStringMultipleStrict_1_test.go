package docker

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetStringMultipleStrict_1(t *testing.T) {
	testCases := []struct {
		name        string
		labels      map[string]string
		labelNames  []string
		expected    map[string]string
		expectedErr error
	}{
		{
			name: "Test case 1",
			labels: map[string]string{
				"label1": "value1",
				"label2": "value2",
				"label3": "value3",
			},
			labelNames: []string{"label1", "label2"},
			expected: map[string]string{
				"label1": "value1",
				"label2": "value2",
			},
			expectedErr: nil,
		},
		{
			name: "Test case 2",
			labels: map[string]string{
				"label1": "value1",
				"label2": "value2",
				"label3": "value3",
			},
			labelNames:  []string{"label1", "label4"},
			expected:    nil,
			expectedErr: fmt.Errorf("label not found: label4"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result, err := getStringMultipleStrict(tc.labels, tc.labelNames...)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

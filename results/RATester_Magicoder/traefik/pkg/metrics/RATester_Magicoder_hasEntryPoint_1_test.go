package metrics

import (
	"fmt"
	"testing"
)

func TestHasEntryPoint_1(t *testing.T) {
	d := &dynamicConfig{
		entryPoints: map[string]bool{
			"entrypoint1": true,
			"entrypoint2": true,
		},
	}

	tests := []struct {
		name         string
		entrypoint   string
		expectedBool bool
	}{
		{
			name:         "Existing entrypoint",
			entrypoint:   "entrypoint1",
			expectedBool: true,
		},
		{
			name:         "Non-existing entrypoint",
			entrypoint:   "entrypoint3",
			expectedBool: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := d.hasEntryPoint(test.entrypoint)
			if result != test.expectedBool {
				t.Errorf("Expected %v, but got %v", test.expectedBool, result)
			}
		})
	}
}

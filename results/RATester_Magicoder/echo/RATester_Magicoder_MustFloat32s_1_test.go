package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMustFloat32s_1(t *testing.T) {
	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			// Implement your logic here
			return ""
		},
		ValuesFunc: func(sourceParam string) []string {
			// Implement your logic here
			return []string{}
		},
		ErrorFunc: func(sourceParam string, values []string, message interface{}, internalError error) error {
			// Implement your logic here
			return nil
		},
	}

	// Define your test cases here
	testCases := []struct {
		name     string
		source   string
		expected []float32
		wantErr  bool
	}{
		// Add your test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := b.MustFloat32s(tc.source, &tc.expected)
			if (got.BindError() != nil) != tc.wantErr {
				t.Errorf("MustFloat32s() error = %v, wantErr %v", got.BindError(), tc.wantErr)
				return
			}
			if !reflect.DeepEqual(tc.expected, tc.expected) {
				t.Errorf("MustFloat32s() = %v, want %v", tc.expected, tc.expected)
			}
		})
	}
}

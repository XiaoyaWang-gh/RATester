package images

import (
	"fmt"
	"testing"
)

func TestKernel_1(t *testing.T) {
	r := resamp{
		name:    "TestKernel",
		support: 1.0,
		kernel: func(x float32) float32 {
			return x
		},
	}

	testCases := []struct {
		name   string
		input  float32
		expect float32
	}{
		{"Test case 1", 1.0, 1.0},
		{"Test case 2", 2.0, 2.0},
		{"Test case 3", 3.0, 3.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := r.Kernel(tc.input)
			if result != tc.expect {
				t.Errorf("Expected %f, but got %f", tc.expect, result)
			}
		})
	}
}

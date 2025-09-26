package images

import (
	"fmt"
	"testing"
)

func TestKernel_1(t *testing.T) {
	testCases := []struct {
		name    string
		support float32
		kernel  func(float32) float32
		x       float32
		want    float32
	}{
		{
			name:    "Test case 1",
			support: 1.0,
			kernel:  func(x float32) float32 { return x },
			x:       2.0,
			want:    2.0,
		},
		{
			name:    "Test case 2",
			support: 2.0,
			kernel:  func(x float32) float32 { return x * 2 },
			x:       3.0,
			want:    6.0,
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			r := resamp{
				name:    tc.name,
				support: tc.support,
				kernel:  tc.kernel,
			}

			got := r.Kernel(tc.x)
			if got != tc.want {
				t.Errorf("Expected %v, but got %v", tc.want, got)
			}
		})
	}
}

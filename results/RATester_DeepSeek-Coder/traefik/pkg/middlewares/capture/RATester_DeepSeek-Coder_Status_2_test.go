package capture

import (
	"fmt"
	"testing"
)

func TestStatus_2(t *testing.T) {
	testCases := []struct {
		name   string
		status int
		want   int
	}{
		{
			name:   "TestStatus_200",
			status: 200,
			want:   200,
		},
		{
			name:   "TestStatus_404",
			status: 404,
			want:   404,
		},
		{
			name:   "TestStatus_500",
			status: 500,
			want:   500,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			crw := &captureResponseWriter{
				status: tc.status,
			}

			if got := crw.Status(); got != tc.want {
				t.Errorf("Status() = %v, want %v", got, tc.want)
			}
		})
	}
}

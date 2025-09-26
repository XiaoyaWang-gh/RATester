package task

import (
	"fmt"
	"testing"
)

func TestGracefulShutdown_1(t *testing.T) {
	testCases := []struct {
		name string
		want <-chan struct{}
	}{
		{
			name: "Test case 1",
			want: globalTaskManager.GracefulShutdown(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := GracefulShutdown()
			if got != tc.want {
				t.Errorf("GracefulShutdown() = %v, want %v", got, tc.want)
			}
		})
	}
}

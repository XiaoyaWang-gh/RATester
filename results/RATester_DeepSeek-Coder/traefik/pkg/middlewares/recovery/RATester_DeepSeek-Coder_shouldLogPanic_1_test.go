package recovery

import (
	"fmt"
	"net/http"
	"testing"
)

func TestShouldLogPanic_1(t *testing.T) {

	tests := []struct {
		name       string
		panicValue interface{}
		want       bool
	}{
		{
			name:       "Test case 1",
			panicValue: nil,
			want:       false,
		},
		{
			name:       "Test case 2",
			panicValue: http.ErrAbortHandler,
			want:       false,
		},
		{
			name:       "Test case 3",
			panicValue: "some other value",
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := shouldLogPanic(tt.panicValue); got != tt.want {
				t.Errorf("shouldLogPanic() = %v, want %v", got, tt.want)
			}
		})
	}
}

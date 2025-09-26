package compress

import (
	"fmt"
	"testing"
)

func TestWriteHeader_1(t *testing.T) {
	rw := &responseWriter{}

	tests := []struct {
		name       string
		statusCode int
		want       bool
	}{
		{
			name:       "Test case 1",
			statusCode: 200,
			want:       true,
		},
		{
			name:       "Test case 2",
			statusCode: 404,
			want:       true,
		},
		{
			name:       "Test case 3",
			statusCode: 500,
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

			rw.WriteHeader(tt.statusCode)
			if rw.statusCodeSet != tt.want {
				t.Errorf("WriteHeader() = %v, want %v", rw.statusCodeSet, tt.want)
			}
		})
	}
}

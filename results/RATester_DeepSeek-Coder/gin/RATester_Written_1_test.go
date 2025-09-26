package gin

import (
	"fmt"
	"testing"
)

func TestWritten_1(t *testing.T) {
	testCases := []struct {
		name   string
		writer *responseWriter
		want   bool
	}{
		{
			name: "TestWritten_True",
			writer: &responseWriter{
				size: 1,
			},
			want: true,
		},
		{
			name: "TestWritten_False",
			writer: &responseWriter{
				size: noWritten,
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.writer.Written()
			if got != tc.want {
				t.Errorf("Written() = %v, want %v", got, tc.want)
			}
		})
	}
}

package context

import (
	"fmt"
	"testing"
)

func TestError_4(t *testing.T) {
	tests := []struct {
		name string
		s    StatusCode
		want string
	}{
		{
			name: "Test StatusCode 200",
			s:    200,
			want: "200",
		},
		{
			name: "Test StatusCode 404",
			s:    404,
			want: "404",
		},
		{
			name: "Test StatusCode 500",
			s:    500,
			want: "500",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.s.Error(); got != tt.want {
				t.Errorf("StatusCode.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

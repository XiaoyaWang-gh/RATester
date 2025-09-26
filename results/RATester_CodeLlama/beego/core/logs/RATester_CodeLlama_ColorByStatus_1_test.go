package logs

import (
	"fmt"
	"testing"
)

func TestColorByStatus_1(t *testing.T) {
	tests := []struct {
		name string
		code int
		want string
	}{
		{
			name: "200",
			code: 200,
			want: "green",
		},
		{
			name: "300",
			code: 300,
			want: "white",
		},
		{
			name: "400",
			code: 400,
			want: "yellow",
		},
		{
			name: "500",
			code: 500,
			want: "red",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ColorByStatus(tt.code); got != tt.want {
				t.Errorf("ColorByStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

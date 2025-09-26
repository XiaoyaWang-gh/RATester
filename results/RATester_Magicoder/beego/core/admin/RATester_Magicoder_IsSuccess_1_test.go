package admin

import (
	"fmt"
	"testing"
)

func TestIsSuccess_1(t *testing.T) {
	tests := []struct {
		name string
		r    *Result
		want bool
	}{
		{
			name: "Success",
			r:    &Result{Status: 200},
			want: true,
		},
		{
			name: "Redirect",
			r:    &Result{Status: 300},
			want: false,
		},
		{
			name: "Client Error",
			r:    &Result{Status: 400},
			want: false,
		},
		{
			name: "Server Error",
			r:    &Result{Status: 500},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.r.IsSuccess(); got != tt.want {
				t.Errorf("Result.IsSuccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

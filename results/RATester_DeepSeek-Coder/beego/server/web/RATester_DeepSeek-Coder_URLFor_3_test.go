package web

import (
	"fmt"
	"testing"
)

func TestURLFor_3(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		values   []interface{}
		want     string
	}{
		{
			name:     "TestURLFor",
			endpoint: "/user/:id",
			values:   []interface{}{1},
			want:     "/user/1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := URLFor(tt.endpoint, tt.values...); got != tt.want {
				t.Errorf("URLFor() = %v, want %v", got, tt.want)
			}
		})
	}
}

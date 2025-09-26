package htime

import (
	"fmt"
	"testing"
	"time"
)

func TestNow_1(t *testing.T) {
	tests := []struct {
		name string
		want time.Time
	}{
		{
			name: "Test case 1",
			want: time.Now(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := Now()
			if !got.Equal(tt.want) {
				t.Errorf("Now() = %v, want %v", got, tt.want)
			}
		})
	}
}

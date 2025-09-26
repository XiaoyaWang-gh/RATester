package warpc

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestgetTimer_1(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
		want *time.Timer
	}{
		{
			name: "Test case 1",
			d:    time.Second,
			want: time.NewTimer(time.Second),
		},
		{
			name: "Test case 2",
			d:    2 * time.Second,
			want: time.NewTimer(2 * time.Second),
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getTimer(tt.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTimer() = %v, want %v", got, tt.want)
			}
		})
	}
}

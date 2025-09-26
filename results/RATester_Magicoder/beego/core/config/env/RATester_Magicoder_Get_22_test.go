package env

import (
	"fmt"
	"testing"
)

func TestGet_22(t *testing.T) {
	tests := []struct {
		name   string
		key    string
		defVal string
		want   string
	}{
		{
			name:   "Test case 1",
			key:    "key1",
			defVal: "default1",
			want:   "value1",
		},
		{
			name:   "Test case 2",
			key:    "key2",
			defVal: "default2",
			want:   "value2",
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

			env.Store(tt.key, tt.want)
			if got := Get(tt.key, tt.defVal); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

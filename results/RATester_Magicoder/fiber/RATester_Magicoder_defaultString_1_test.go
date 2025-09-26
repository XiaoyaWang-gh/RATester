package fiber

import (
	"fmt"
	"testing"
)

func TestdefaultString_1(t *testing.T) {
	tests := []struct {
		name         string
		value        string
		defaultValue []string
		want         string
	}{
		{
			name:         "Test case 1",
			value:        "",
			defaultValue: []string{"default"},
			want:         "default",
		},
		{
			name:         "Test case 2",
			value:        "value",
			defaultValue: []string{"default"},
			want:         "value",
		},
		{
			name:         "Test case 3",
			value:        "",
			defaultValue: []string{},
			want:         "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := defaultString(tt.value, tt.defaultValue); got != tt.want {
				t.Errorf("defaultString() = %v, want %v", got, tt.want)
			}
		})
	}
}

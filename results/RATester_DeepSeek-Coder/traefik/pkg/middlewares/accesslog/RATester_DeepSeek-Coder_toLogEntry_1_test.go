package accesslog

import (
	"fmt"
	"testing"
)

func TestToLogEntry_1(t *testing.T) {
	type test struct {
		name         string
		s            string
		defaultValue string
		quote        bool
		want         string
	}

	tests := []test{
		{
			name:         "Test case 1",
			s:            "test",
			defaultValue: "default",
			quote:        true,
			want:         `"test"`,
		},
		{
			name:         "Test case 2",
			s:            "",
			defaultValue: "default",
			quote:        true,
			want:         "default",
		},
		{
			name:         "Test case 3",
			s:            "test",
			defaultValue: "default",
			quote:        false,
			want:         "test",
		},
		{
			name:         "Test case 4",
			s:            "",
			defaultValue: "default",
			quote:        false,
			want:         "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := toLogEntry(tt.s, tt.defaultValue, tt.quote); got != tt.want {
				t.Errorf("toLogEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}

package hugo

import (
	"fmt"
	"testing"
)

func TestCompare_1(t *testing.T) {
	tests := []struct {
		name  string
		h     VersionString
		other any
		want  int
	}{
		{
			name:  "equal",
			h:     "0.0.0",
			other: "0.0.0",
			want:  0,
		},
		{
			name:  "less",
			h:     "0.0.0",
			other: "0.0.1",
			want:  -1,
		},
		{
			name:  "greater",
			h:     "0.0.1",
			other: "0.0.0",
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.h.Compare(tt.other); got != tt.want {
				t.Errorf("VersionString.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

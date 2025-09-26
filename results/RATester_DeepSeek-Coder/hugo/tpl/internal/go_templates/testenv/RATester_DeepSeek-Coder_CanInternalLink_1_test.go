package testenv

import (
	"fmt"
	"testing"
)

func TestCanInternalLink_1(t *testing.T) {
	tests := []struct {
		name    string
		withCgo bool
		want    bool
	}{
		{
			name:    "Test case 1",
			withCgo: false,
			want:    false,
		},
		{
			name:    "Test case 2",
			withCgo: true,
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := CanInternalLink(tt.withCgo); got != tt.want {
				t.Errorf("CanInternalLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

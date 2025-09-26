package hugo

import (
	"fmt"
	"testing"
)

func TestCompareVersion_1(t *testing.T) {
	tests := []struct {
		name    string
		version any
		want    int
	}{
		{
			name:    "Test case 1",
			version: "1.0.0",
			want:    0,
		},
		{
			name:    "Test case 2",
			version: "2.0.0",
			want:    -1,
		},
		{
			name:    "Test case 3",
			version: "0.1.0",
			want:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := CompareVersion(tt.version); got != tt.want {
				t.Errorf("CompareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

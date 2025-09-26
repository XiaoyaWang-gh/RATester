package hugo

import (
	"fmt"
	"testing"
)

func TestVersion_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var tests = []struct {
		major, minor, patch int
		suffix              string
		want                string
	}{
		{1, 2, 3, "test", "1.2.3test"},
		{1, 2, 0, "test", "1.2test"},
		{1, 0, 0, "test", "1test"},
		{1, 2, 3, "", "1.2.3"},
		{1, 2, 0, "", "1.2"},
		{1, 0, 0, "", "1"},
	}
	for _, tt := range tests {
		if got := version(tt.major, tt.minor, tt.patch, tt.suffix); got != tt.want {
			t.Errorf("version(%d, %d, %d, %q) = %q; want %q", tt.major, tt.minor, tt.patch, tt.suffix, got, tt.want)
		}
	}
}

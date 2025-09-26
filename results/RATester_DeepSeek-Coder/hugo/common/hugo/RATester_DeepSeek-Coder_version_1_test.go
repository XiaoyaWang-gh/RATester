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

	type test struct {
		major, minor, patch int
		suffix              string
		expected            string
	}

	tests := []test{
		{1, 0, 0, "", "1.0"},
		{1, 2, 3, "-beta", "1.2-beta"},
		{1, 2, 3, "-rc1", "1.2.3-rc1"},
		{1, 2, 0, "-rc1", "1.2-rc1"},
		{1, 53, 0, "-rc1", "1.53-rc1"},
		{1, 54, 0, "-rc1", "1.54.0-rc1"},
		{1, 53, 1, "-rc1", "1.53.1-rc1"},
	}

	for _, test := range tests {
		result := version(test.major, test.minor, test.patch, test.suffix)
		if result != test.expected {
			t.Errorf("version(%d, %d, %d, %s) = %s; want %s", test.major, test.minor, test.patch, test.suffix, result, test.expected)
		}
	}
}

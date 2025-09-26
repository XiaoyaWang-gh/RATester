package cors

import (
	"fmt"
	"testing"
)

func TestIsOriginAllowed_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	o := &Options{
		AllowAllOrigins: true,
	}

	tests := []struct {
		origin  string
		allowed bool
	}{
		{"http://example.com", true},
		{"http://notallowed.com", false},
	}

	for _, test := range tests {
		allowed := o.IsOriginAllowed(test.origin)
		if allowed != test.allowed {
			t.Errorf("IsOriginAllowed(%s) = %v, want %v", test.origin, allowed, test.allowed)
		}
	}
}

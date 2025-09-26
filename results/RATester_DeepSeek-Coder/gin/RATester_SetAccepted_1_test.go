package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetAccepted_1(t *testing.T) {
	tests := []struct {
		name     string
		formats  []string
		expected []string
	}{
		{
			name:     "TestSetAccepted",
			formats:  []string{"application/json", "application/xml"},
			expected: []string{"application/json", "application/xml"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := &Context{}
			ctx.SetAccepted(test.formats...)

			if !reflect.DeepEqual(ctx.Accepted, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, ctx.Accepted)
			}
		})
	}
}

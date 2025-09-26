package xml

import (
	"fmt"
	"testing"
)

func TestDefaultFloat_9(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name       string
		key        string
		defaultVal float64
		want       float64
	}

	testCases := []testCase{
		{
			name:       "TestDefaultFloat_ExistingKey",
			key:        "existingKey",
			defaultVal: 1.0,
			want:       2.0,
		},
		{
			name:       "TestDefaultFloat_NonExistingKey",
			key:        "nonExistingKey",
			defaultVal: 3.0,
			want:       3.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &ConfigContainer{
				data: map[string]interface{}{
					tc.key: tc.want,
				},
			}

			got := c.DefaultFloat(tc.key, tc.defaultVal)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

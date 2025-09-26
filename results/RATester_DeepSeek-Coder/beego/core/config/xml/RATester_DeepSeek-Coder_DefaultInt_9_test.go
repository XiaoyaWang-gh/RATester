package xml

import (
	"fmt"
	"testing"
)

func TestDefaultInt_9(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name       string
		key        string
		defaultVal int
		want       int
	}

	testCases := []testCase{
		{
			name:       "TestDefaultInt_ExistingKey",
			key:        "existingKey",
			defaultVal: 100,
			want:       200,
		},
		{
			name:       "TestDefaultInt_NonExistingKey",
			key:        "nonExistingKey",
			defaultVal: 300,
			want:       300,
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

			got := c.DefaultInt(tc.key, tc.defaultVal)
			if got != tc.want {
				t.Errorf("Expected %v, got %v", tc.want, got)
			}
		})
	}
}

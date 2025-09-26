package config

import (
	"fmt"
	"strconv"
	"testing"
)

func TestDefaultInt_3(t *testing.T) {
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
			defaultVal: 100,
			want:       100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &fakeConfigContainer{
				data: map[string]string{
					tc.key: strconv.Itoa(tc.want),
				},
			}

			got := c.DefaultInt(tc.key, tc.defaultVal)
			if got != tc.want {
				t.Errorf("DefaultInt() = %v, want %v", got, tc.want)
			}
		})
	}
}

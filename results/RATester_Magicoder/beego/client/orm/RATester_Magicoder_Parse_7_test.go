package orm

import (
	"fmt"
	"testing"
)

func TestParse_7(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expected *alias
	}{
		{
			name: "Test with default db",
			args: []string{},
			expected: &alias{
				Name: "default",
			},
		},
		{
			name: "Test with custom db",
			args: []string{"-db", "custom"},
			expected: &alias{
				Name: "custom",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := &commandSQLAll{}
			d.Parse(tc.args)

			if d.al.Name != tc.expected.Name {
				t.Errorf("Expected Name: %s, but got: %s", tc.expected.Name, d.al.Name)
			}
		})
	}
}

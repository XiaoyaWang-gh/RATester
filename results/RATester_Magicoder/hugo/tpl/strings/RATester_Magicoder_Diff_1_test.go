package strings

import (
	"fmt"
	"testing"
)

func TestDiff_1(t *testing.T) {
	ns := &Namespace{}
	tests := []struct {
		name     string
		oldname  string
		old      any
		newname  string
		new      any
		expected string
		err      error
	}{
		{
			name:     "Test case 1",
			oldname:  "old",
			old:      "old value",
			newname:  "new",
			new:      "new value",
			expected: "diff.Diff(oldname, []byte(olds), newname, []byte(news))",
			err:      nil,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.Diff(tt.oldname, tt.old, tt.newname, tt.new)
			if err != tt.err {
				t.Errorf("Diff() error = %v, wantErr %v", err, tt.err)
				return
			}
			if got != tt.expected {
				t.Errorf("Diff() = %v, want %v", got, tt.expected)
			}
		})
	}
}

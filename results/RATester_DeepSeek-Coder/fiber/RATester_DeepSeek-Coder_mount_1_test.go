package fiber

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestMount_1(t *testing.T) {
	app := New()
	subApp := New()
	grp := &Group{
		app: app,
	}

	tests := []struct {
		name     string
		grp      *Group
		prefix   string
		subApp   *App
		expected *Group
	}{
		{
			name:   "Test mount with empty prefix",
			grp:    grp,
			prefix: "",
			subApp: subApp,
			expected: &Group{
				Prefix: "/",
				app:    subApp,
			},
		},
		{
			name:   "Test mount with non-empty prefix",
			grp:    grp,
			prefix: "/api",
			subApp: subApp,
			expected: &Group{
				Prefix: "/api",
				app:    subApp,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tt.grp.mount(tt.prefix, tt.subApp)
			assert.Equal(t, tt.expected, result)
		})
	}
}

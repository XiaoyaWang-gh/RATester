package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func Testmount_1(t *testing.T) {
	app := &App{}
	subApp := &App{}
	group := &Group{Prefix: "/", app: app}

	tests := []struct {
		name   string
		grp    *Group
		prefix string
		subApp *App
		want   Router
	}{
		{
			name:   "Test Case 1",
			grp:    group,
			prefix: "/test",
			subApp: subApp,
			want:   group,
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

			if got := tt.grp.mount(tt.prefix, tt.subApp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mount() = %v, want %v", got, tt.want)
			}
		})
	}
}

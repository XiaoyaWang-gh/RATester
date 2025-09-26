package web

import (
	"fmt"
	"testing"
)

func TestURLFor_1(t *testing.T) {
	ctrlRegister := &ControllerRegister{
		routers: make(map[string]*Tree),
	}

	ctrlRegister.routers["GET"] = &Tree{}

	testCases := []struct {
		name     string
		endpoint string
		values   []interface{}
		want     string
	}{
		{
			name:     "TestURLFor_0",
			endpoint: "path.controller.method",
			values:   []interface{}{"key", "value"},
			want:     "/path/controller/method?key=value",
		},
		{
			name:     "TestURLFor_1",
			endpoint: "path.controller",
			values:   []interface{}{"key", "value"},
			want:     "",
		},
		{
			name:     "TestURLFor_2",
			endpoint: "path",
			values:   []interface{}{"key", "value"},
			want:     "",
		},
		{
			name:     "TestURLFor_3",
			endpoint: "path.controller.method",
			values:   []interface{}{},
			want:     "/path/controller/method",
		},
		{
			name:     "TestURLFor_4",
			endpoint: "path.controller.method",
			values:   []interface{}{"key"},
			want:     "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := ctrlRegister.URLFor(tc.endpoint, tc.values...)
			if got != tc.want {
				t.Errorf("URLFor() = %v, want %v", got, tc.want)
			}
		})
	}
}

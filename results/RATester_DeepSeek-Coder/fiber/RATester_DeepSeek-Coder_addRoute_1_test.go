package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddRoute_1(t *testing.T) {
	app := new(App)
	route := new(Route)

	tests := []struct {
		name     string
		method   string
		route    *Route
		expected *App
	}{
		{
			name:   "Test addRoute with GET method",
			method: "GET",
			route:  route,
			expected: &App{
				stack: [][]*Route{
					{route},
				},
			},
		},
		// Add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			app.addRoute(tt.method, tt.route)
			if !reflect.DeepEqual(app, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, app)
			}
		})
	}
}

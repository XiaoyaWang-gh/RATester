package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestSetRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TEST CASE SETUP
	// ### important: always keep in sync with the copy method "app.copyRoute" ###
	route := &Route{
		group: &Group{
			app: &App{
				config: Config{
					StrictRouting: true,
				},
			},
		},
		path: "/foo",
	}
	// ### important: always keep in sync with the copy method "app.copyRoute" ###
	c := &DefaultCtx{
		route: route,
	}
	// TEST CASE EXECUTION
	c.setRoute(route)
	// ASSERTIONS
	assert.Equal(t, route, c.route)
	assert.Equal(t, route.path, c.path)
	assert.Equal(t, route.path, c.detectionPath)
}

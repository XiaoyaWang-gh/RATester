package fiber

import (
	"fmt"
	"testing"
)

func TestgetLocationFromRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		app: &App{
			config: Config{
				CaseSensitive: false,
			},
		},
	}

	route := &Route{
		routeParser: routeParser{
			segs: []*routeSegment{
				{
					Const: "const",
				},
				{
					Const:       "param",
					ParamName:   "param",
					IsParam:     true,
					IsGreedy:    true,
					IsOptional:  true,
					ComparePart: "param",
				},
			},
		},
	}

	params := Map{
		"param": "value",
	}

	location, err := ctx.getLocationFromRoute(*route, params)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := "constvalue"
	if location != expected {
		t.Errorf("Expected %s, but got %s", expected, location)
	}
}

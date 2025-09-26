package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuildTree_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		config: Config{
			RequestMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS", "CONNECT", "TRACE"},
		},
		stack: [][]*Route{
			{
				{
					path: "/",
					routeParser: routeParser{
						segs: []*routeSegment{
							{
								Const:    "/",
								IsLast:   true,
								IsGreedy: false,
							},
						},
					},
				},
			},
			{
				{
					path: "/test",
					routeParser: routeParser{
						segs: []*routeSegment{
							{
								Const:    "/",
								IsLast:   true,
								IsGreedy: false,
							},
						},
					},
				},
			},
		},
	}

	expectedTreeStack := []map[string][]*Route{
		{
			"": {
				{
					path: "/",
					routeParser: routeParser{
						segs: []*routeSegment{
							{
								Const:    "/",
								IsLast:   true,
								IsGreedy: false,
							},
						},
					},
				},
			},
		},
		{
			"": {
				{
					path: "/test",
					routeParser: routeParser{
						segs: []*routeSegment{
							{
								Const:    "/",
								IsLast:   true,
								IsGreedy: false,
							},
						},
					},
				},
			},
		},
	}

	app.buildTree()

	if !reflect.DeepEqual(app.treeStack, expectedTreeStack) {
		t.Errorf("Expected treeStack to be %v, but got %v", expectedTreeStack, app.treeStack)
	}
}

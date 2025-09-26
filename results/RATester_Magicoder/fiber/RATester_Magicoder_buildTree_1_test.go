package fiber

import (
	"fmt"
	"testing"
)

func TestbuildTree_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		routesRefreshed: true,
		config: Config{
			RequestMethods: []string{"GET", "POST"},
		},
		stack: [][]*Route{
			{
				{
					routeParser: routeParser{
						segs: []*routeSegment{
							{
								Const: "abc",
							},
						},
					},
				},
			},
		},
		treeStack: []map[string][]*Route{
			{
				"": {
					{
						routeParser: routeParser{
							segs: []*routeSegment{
								{
									Const: "abc",
								},
							},
						},
					},
				},
			},
		},
	}

	app.buildTree()

	if app.routesRefreshed {
		t.Error("routesRefreshed should be false after buildTree")
	}

	if len(app.treeStack[0]) != 1 {
		t.Error("treeStack should have 1 element")
	}

	if len(app.treeStack[0]["abc"]) != 1 {
		t.Error("treeStack should have 1 route")
	}
}

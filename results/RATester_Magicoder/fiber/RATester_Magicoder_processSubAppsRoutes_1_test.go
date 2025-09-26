package fiber

import (
	"fmt"
	"testing"
)

func TestprocessSubAppsRoutes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		mountFields: &mountFields{
			appList: map[string]*App{
				"": &App{
					stack: [][]*Route{
						{
							{
								group: &Group{
									app: &App{
										stack: [][]*Route{
											{
												{
													path: "/subapp",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	app.processSubAppsRoutes()

	// Add assertions here
}

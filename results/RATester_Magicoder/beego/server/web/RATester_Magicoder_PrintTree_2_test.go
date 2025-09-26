package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrintTree_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{
		Handlers: &ControllerRegister{
			routers: map[string]*Tree{
				"GET": {
					prefix: "/api",
					leaves: []*leafInfo{
						{
							wildcards: []string{"id"},
							runObject: "GET /api/:id",
						},
					},
				},
				"POST": {
					prefix: "/api",
					leaves: []*leafInfo{
						{
							wildcards: []string{"id"},
							runObject: "POST /api/:id",
						},
					},
				},
			},
		},
	}

	expected := M{
		"Data": M{
			"GET":  [][]string{{"GET /api/:id"}},
			"POST": [][]string{{"POST /api/:id"}},
		},
		"Methods": []string{"GET", "POST"},
	}

	result := app.PrintTree()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

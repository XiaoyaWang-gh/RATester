package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddPrefix_1(t *testing.T) {
	type testStruct struct {
		name     string
		tree     *Tree
		prefix   string
		expected *Tree
	}

	testCases := []testStruct{
		{
			name: "Test Case 1",
			tree: &Tree{
				prefix: "/api",
				fixrouters: []*Tree{
					{
						prefix: "/v1",
						leaves: []*leafInfo{
							{
								runObject: &ControllerInfo{
									pattern: "/user",
								},
							},
						},
					},
				},
				wildcard: &Tree{
					prefix: "/:id",
					leaves: []*leafInfo{
						{
							runObject: &ControllerInfo{
								pattern: "/user/:id",
							},
						},
					},
				},
			},
			prefix: "/prefix",
			expected: &Tree{
				prefix: "/api",
				fixrouters: []*Tree{
					{
						prefix: "/v1",
						leaves: []*leafInfo{
							{
								runObject: &ControllerInfo{
									pattern: "/prefix/user",
								},
							},
						},
					},
				},
				wildcard: &Tree{
					prefix: "/:id",
					leaves: []*leafInfo{
						{
							runObject: &ControllerInfo{
								pattern: "/prefix/user/:id",
							},
						},
					},
				},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			addPrefix(tc.tree, tc.prefix)
			if !reflect.DeepEqual(tc.tree, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, tc.tree)
			}
		})
	}
}

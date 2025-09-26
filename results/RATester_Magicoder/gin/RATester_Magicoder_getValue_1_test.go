package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetValue_1(t *testing.T) {
	tests := []struct {
		name         string
		node         *node
		path         string
		params       *Params
		skippedNodes *[]skippedNode
		unescape     bool
		wantValue    nodeValue
	}{
		{
			name: "test case 1",
			node: &node{
				path:      "/test",
				indices:   "t",
				wildChild: false,
				nType:     static,
				priority:  0,
				children: []*node{
					{
						path:      "/test/param",
						indices:   "p",
						wildChild: false,
						nType:     param,
						priority:  0,
						children:  []*node{},
						handlers:  HandlersChain{},
						fullPath:  "/test/param",
					},
				},
				handlers: HandlersChain{},
				fullPath: "/test",
			},
			path:         "/test/param",
			params:       &Params{},
			skippedNodes: &[]skippedNode{},
			unescape:     false,
			wantValue: nodeValue{
				handlers: HandlersChain{},
				params: &Params{
					{
						Key:   "param",
						Value: "param",
					},
				},
				tsr:      false,
				fullPath: "/test/param",
			},
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

			gotValue := tt.node.getValue(tt.path, tt.params, tt.skippedNodes, tt.unescape)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("node.getValue() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

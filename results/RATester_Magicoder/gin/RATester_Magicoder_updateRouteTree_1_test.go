package gin

import (
	"fmt"
	"testing"
)

func TestupdateRouteTree_1(t *testing.T) {
	type args struct {
		n *node
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			updateRouteTree(tt.args.n)
		})
	}
}

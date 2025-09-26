package web

import (
	"fmt"
	"testing"
)

func TestfindAndRemoveTree_1(t *testing.T) {
	type args struct {
		paths          []string
		entryPointTree *Tree
		method         string
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

			findAndRemoveTree(tt.args.paths, tt.args.entryPointTree, tt.args.method)
		})
	}
}

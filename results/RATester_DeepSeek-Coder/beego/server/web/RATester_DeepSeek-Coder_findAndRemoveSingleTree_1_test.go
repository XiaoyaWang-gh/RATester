package web

import (
	"fmt"
	"testing"
)

func TestFindAndRemoveSingleTree_1(t *testing.T) {
	type args struct {
		entryPointTree *Tree
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

			findAndRemoveSingleTree(tt.args.entryPointTree)
		})
	}
}

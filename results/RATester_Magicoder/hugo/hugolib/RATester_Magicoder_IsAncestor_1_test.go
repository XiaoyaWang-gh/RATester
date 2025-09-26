package hugolib

import (
	"fmt"
	"testing"
)

func TestIsAncestor_1(t *testing.T) {
	type args struct {
		other any
	}
	tests := []struct {
		name string
		pt   pageTree
		args args
		want bool
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

			if got := tt.pt.IsAncestor(tt.args.other); got != tt.want {
				t.Errorf("pageTree.IsAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}

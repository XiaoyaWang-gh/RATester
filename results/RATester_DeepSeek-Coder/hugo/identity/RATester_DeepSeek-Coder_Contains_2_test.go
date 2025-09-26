package identity

import (
	"fmt"
	"testing"
)

func TestContains_2(t *testing.T) {
	type args struct {
		id       Identity
		in       Identity
		maxDepth int
	}
	tests := []struct {
		name string
		args args
		want FinderResult
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

			f := &Finder{}
			if got := f.Contains(tt.args.id, tt.args.in, tt.args.maxDepth); got != tt.want {
				t.Errorf("Finder.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

package cli

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/paerser/parser"
)

func TestHasKnownNodes_1(t *testing.T) {
	type args struct {
		rootType reflect.Type
		node     *parser.Node
	}
	tests := []struct {
		name string
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

			if got := hasKnownNodes(tt.args.rootType, tt.args.node); got != tt.want {
				t.Errorf("hasKnownNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

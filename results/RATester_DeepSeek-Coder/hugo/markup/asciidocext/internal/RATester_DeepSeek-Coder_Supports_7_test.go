package internal

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestSupports_7(t *testing.T) {
	type args struct {
		identity identity.Identity
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

			a := &AsciidocConverter{}
			if got := a.Supports(tt.args.identity); got != tt.want {
				t.Errorf("AsciidocConverter.Supports() = %v, want %v", got, tt.want)
			}
		})
	}
}

package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestForEeachIdentity_2(t *testing.T) {
	type args struct {
		f func(identity.Identity) bool
	}
	tests := []struct {
		name string
		p    *pageState
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

			if got := tt.p.ForEeachIdentity(tt.args.f); got != tt.want {
				t.Errorf("ForEeachIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}

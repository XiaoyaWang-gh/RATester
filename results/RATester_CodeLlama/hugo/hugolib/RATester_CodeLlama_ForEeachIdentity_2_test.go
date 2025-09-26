package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestForEeachIdentity_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	var p *pageState
	var f func(identity.Identity) bool

	// When
	p.ForEeachIdentity(f)

	// Then
	// ...
}

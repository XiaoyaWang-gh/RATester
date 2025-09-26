package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestForEeachIdentityByName_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	var name string
	var f func(identity.Identity) bool
	var s *Site
	// When
	s.ForEeachIdentityByName(name, f)
	// Then
	// ...
}

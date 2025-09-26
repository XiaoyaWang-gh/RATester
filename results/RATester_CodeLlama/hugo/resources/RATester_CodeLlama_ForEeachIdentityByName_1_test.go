package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestForEeachIdentityByName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	name := "name"
	f := func(identity.Identity) bool {
		return true
	}
	r := &resourceAdapter{}

	// when
	r.ForEeachIdentityByName(name, f)

	// then
	// TODO: add assertions
}

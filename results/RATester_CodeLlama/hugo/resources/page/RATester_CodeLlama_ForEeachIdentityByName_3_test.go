package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestForEeachIdentityByName_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	s := &siteWrapper{}
	name := "name"
	f := func(identity.Identity) bool {
		return true
	}

	// when
	s.ForEeachIdentityByName(name, f)

	// then
	// ...
}

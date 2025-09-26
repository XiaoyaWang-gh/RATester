package pagesfromdata

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestAddChange_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var b PagesFromTemplate
	var id identity.Identity
	b.AddChange(id)
	if len(b.buildState.ChangedIdentities) != 1 {
		t.Errorf("Expected 1 changed identity, got %d", len(b.buildState.ChangedIdentities))
	}
}

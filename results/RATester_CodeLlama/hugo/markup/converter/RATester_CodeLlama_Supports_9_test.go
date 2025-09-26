package converter

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestSupports_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var feature identity.Identity
	var nopConverter nopConverter
	if nopConverter.Supports(feature) {
		t.Errorf("nopConverter.Supports(feature) = true, want false")
	}
}

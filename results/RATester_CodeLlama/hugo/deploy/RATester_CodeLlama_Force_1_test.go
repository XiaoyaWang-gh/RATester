package deploy

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
)

func TestForce_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lf := &localFile{matcher: &deployconfig.Matcher{Force: true}}
	if !lf.Force() {
		t.Errorf("Force() = false, want true")
	}
}

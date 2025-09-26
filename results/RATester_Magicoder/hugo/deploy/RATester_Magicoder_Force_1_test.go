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

	lf := &localFile{
		matcher: &deployconfig.Matcher{
			Force: true,
		},
	}

	if !lf.Force() {
		t.Error("Expected Force to return true")
	}

	lf.matcher = nil
	if lf.Force() {
		t.Error("Expected Force to return false when matcher is nil")
	}

	lf.matcher = &deployconfig.Matcher{
		Force: false,
	}
	if lf.Force() {
		t.Error("Expected Force to return false when matcher.Force is false")
	}
}

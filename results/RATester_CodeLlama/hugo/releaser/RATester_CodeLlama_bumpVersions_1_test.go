package releaser

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/hugo"
)

func TestBumpVersions_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &ReleaseHandler{}
	ver := hugo.Version{
		Major:      0,
		Minor:      90,
		PatchLevel: 0,
		Suffix:     "",
	}
	err := r.bumpVersions(ver)
	if err != nil {
		t.Errorf("bumpVersions() error = %v", err)
		return
	}
}

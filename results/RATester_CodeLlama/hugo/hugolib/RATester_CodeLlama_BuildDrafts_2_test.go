package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config/allconfig"
)

func TestBuildDrafts_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	s := &Site{}
	if s.BuildDrafts() {
		t.Error("BuildDrafts() should be false by default")
	}
	s.conf = &allconfig.Config{}
	if !s.BuildDrafts() {
		t.Error("BuildDrafts() should be true if conf.BuildDrafts is true")
	}
}

package js

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/gohugoio/hugo/resources"
)

func TestNew_26(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &filesystems.SourceFilesystem{}
	rs := &resources.Spec{}
	c := New(fs, rs)
	if c.rs != rs {
		t.Errorf("rs not set")
	}
	if c.sfs != fs {
		t.Errorf("fs not set")
	}
}

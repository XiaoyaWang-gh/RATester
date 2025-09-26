package scss

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugolib/filesystems"
	"github.com/gohugoio/hugo/resources"
)

func TestNew_42(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// CONTEXT_0
	fs := &filesystems.SourceFilesystem{}
	// CONTEXT_1
	rs := &resources.Spec{}
	// CONTEXT_2
	c, err := New(fs, rs)
	if err != nil {
		t.Fatal(err)
	}
	// CONTEXT_3
	if c.rs != rs {
		t.Errorf("expected %v, got %v", rs, c.rs)
	}
	// CONTEXT_4
	if c.sfs != fs {
		t.Errorf("expected %v, got %v", fs, c.sfs)
	}
	// CONTEXT_5
	if c.workFs != rs.BaseFs.Work {
		t.Errorf("expected %v, got %v", rs.BaseFs.Work, c.workFs)
	}
	// CONTEXT_6
	if _, err := c.ToCSS(nil, Options{}); err != nil {
		t.Error(err)
	}
	// CONTEXT_7
	if err != nil {
		t.Error(err)
	}
	// CONTEXT_8
	if c.sfs != fs {
		t.Errorf("expected %v, got %v", fs, c.sfs)
	}
	// CONTEXT_9
	if c.workFs != rs.BaseFs.Work {
		t.Errorf("expected %v, got %v", rs.BaseFs.Work, c.workFs)
	}
	// CONTEXT_10
	if _, err := c.ToCSS(nil, Options{}); err != nil {
		t.Error(err)
	}
	// CONTEXT_11
	if err != nil {
		t.Error(err)
	}
}

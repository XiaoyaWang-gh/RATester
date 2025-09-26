package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/paths"
	"github.com/gohugoio/hugo/helpers"
	"github.com/gohugoio/hugo/output"
)

func TestPutPagePathBuilder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &pagePathBuilder{
		els: []string{"test1", "test2"},
		d: TargetPathDescriptor{
			PathSpec:          &helpers.PathSpec{},
			Type:              output.Format{},
			Kind:              "testKind",
			Path:              &paths.Path{},
			Section:           &paths.Path{},
			BaseName:          "testBaseName",
			PrefixFilePath:    "testPrefixFilePath",
			PrefixLink:        "testPrefixLink",
			ForcePrefix:       true,
			URL:               "testURL",
			Addends:           "testAddends",
			ExpandedPermalink: "testExpandedPermalink",
			UglyURLs:          true,
		},
		isUgly:             true,
		baseNameSameAsType: true,
	}

	putPagePathBuilder(b)

	if len(b.els) != 0 {
		t.Errorf("Expected els to be empty, got %v", b.els)
	}

	if b.fullSuffix != "" {
		t.Errorf("Expected fullSuffix to be empty, got %v", b.fullSuffix)
	}

	if b.baseNameSameAsType != false {
		t.Errorf("Expected baseNameSameAsType to be false, got %v", b.baseNameSameAsType)
	}

	if b.isUgly != false {
		t.Errorf("Expected isUgly to be false, got %v", b.isUgly)
	}

	if b.prefixLink != "" {
		t.Errorf("Expected prefixLink to be empty, got %v", b.prefixLink)
	}

	if b.prefixPath != "" {
		t.Errorf("Expected prefixPath to be empty, got %v", b.prefixPath)
	}

	if b.linkUpperOffset != 0 {
		t.Errorf("Expected linkUpperOffset to be 0, got %v", b.linkUpperOffset)
	}
}

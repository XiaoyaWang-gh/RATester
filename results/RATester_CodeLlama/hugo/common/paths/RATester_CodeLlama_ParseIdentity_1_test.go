package paths

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestParseIdentity_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pp := &PathParser{
		LanguageIndex: map[string]int{
			"en": 0,
		},
		IsLangDisabled: func(lang string) bool {
			return false
		},
		IsContentExt: func(ext string) bool {
			return false
		},
	}

	tests := []struct {
		c, s string
		want identity.StringIdentity
	}{
		{"en", "foo", "foo"},
		{"en", "foo/bar", "foo/bar"},
		{"en", "foo/bar/", "foo/bar"},
		{"en", "foo/bar/baz", "foo/bar/baz"},
		{"en", "foo/bar/baz/", "foo/bar/baz"},
		{"en", "foo/bar/baz/qux", "foo/bar/baz/qux"},
		{"en", "foo/bar/baz/qux/", "foo/bar/baz/qux"},
	}

	for _, test := range tests {
		if got := pp.ParseIdentity(test.c, test.s); got != test.want {
			t.Errorf("ParseIdentity(%q, %q) = %q, want %q", test.c, test.s, got, test.want)
		}
	}
}

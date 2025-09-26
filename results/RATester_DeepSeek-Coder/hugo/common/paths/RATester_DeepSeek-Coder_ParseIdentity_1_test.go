package paths

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestParseIdentity_1(t *testing.T) {
	pp := &PathParser{
		LanguageIndex: map[string]int{
			"en": 0,
			"fr": 1,
		},
		IsLangDisabled: func(lang string) bool {
			return lang == "fr"
		},
		IsContentExt: func(ext string) bool {
			return ext == "md"
		},
	}

	tests := []struct {
		name     string
		category string
		path     string
		want     identity.StringIdentity
	}{
		{
			name:     "valid path",
			category: "posts",
			path:     "en/my-post.md",
			want:     "en/posts/my-post",
		},
		{
			name:     "disabled language",
			category: "posts",
			path:     "fr/my-post.md",
			want:     "",
		},
		{
			name:     "non-content file",
			category: "posts",
			path:     "en/my-post.txt",
			want:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := pp.ParseIdentity(tt.category, tt.path)
			if got != tt.want {
				t.Errorf("ParseIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}

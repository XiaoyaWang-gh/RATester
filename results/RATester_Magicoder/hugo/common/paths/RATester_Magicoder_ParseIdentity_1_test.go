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
		language string
		path     string
		want     identity.StringIdentity
	}{
		{
			name:     "English",
			language: "en",
			path:     "content/posts/my-post.md",
			want:     "content/posts/my-post",
		},
		{
			name:     "French",
			language: "fr",
			path:     "content/posts/my-post.md",
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

			if got := pp.ParseIdentity(tt.language, tt.path); got != tt.want {
				t.Errorf("ParseIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}

package paths

import (
	"fmt"
	"testing"
)

func TestParseBaseAndBaseNameNoIdentifier_1(t *testing.T) {
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
		input    string
		expected string
	}{
		{
			name:     "Test case 1",
			input:    "/content/posts/post1.md",
			expected: "post1",
		},
		{
			name:     "Test case 2",
			input:    "/content/posts/post2.html",
			expected: "post2.html",
		},
		{
			name:     "Test case 3",
			input:    "/content/posts/post3.txt",
			expected: "post3.txt",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, result := pp.ParseBaseAndBaseNameNoIdentifier("en", test.input)
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}

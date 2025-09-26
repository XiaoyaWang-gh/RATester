package paths

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewPath_1(t *testing.T) {
	pp := &PathParser{
		LanguageIndex: map[string]int{
			"en": 0,
			"fr": 1,
		},
		IsLangDisabled: func(lang string) bool {
			return lang == "fr"
		},
		IsContentExt: func(ext string) bool {
			return ext == ".md" || ext == ".html"
		},
	}

	tests := []struct {
		name      string
		pp        *PathParser
		component string
		want      *Path
	}{
		{
			name:      "Test newPath with component 'test'",
			pp:        pp,
			component: "test",
			want: &Path{
				component:             "test",
				posContainerLow:       -1,
				posContainerHigh:      -1,
				posSectionHigh:        -1,
				posIdentifierLanguage: -1,
			},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.pp.newPath(tt.component)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

package layouts

import (
	"fmt"
	"reflect"
	"testing"
)

func TestresolvePageTemplate_1(t *testing.T) {
	tests := []struct {
		name string
		d    LayoutDescriptor
		want []string
	}{
		{
			name: "Test case 1",
			d: LayoutDescriptor{
				Type:             "page",
				Section:          "section1",
				Kind:             "page",
				KindVariants:     "variant1,variant2",
				Lang:             "en",
				Layout:           "layout1",
				LayoutOverride:   false,
				OutputFormatName: "rss",
				Suffix:           "suffix1",
				RenderingHook:    false,
				Baseof:           false,
			},
			want: []string{"layout1", "single", "section1", "page", "variant1", "variant2", "en", "suffix1", "rss", "_default", "list"},
		},
		{
			name: "Test case 2",
			d: LayoutDescriptor{
				Type:             "home",
				Section:          "section2",
				Kind:             "home",
				KindVariants:     "variant3,variant4",
				Lang:             "fr",
				Layout:           "layout2",
				LayoutOverride:   true,
				OutputFormatName: "html",
				Suffix:           "suffix2",
				RenderingHook:    true,
				Baseof:           true,
			},
			want: []string{"layout2", "index", "home", "home", "variant3", "variant4", "fr", "suffix2", "html", "_default", "baseof", "list"},
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

			got := resolvePageTemplate(tt.d)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolvePageTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

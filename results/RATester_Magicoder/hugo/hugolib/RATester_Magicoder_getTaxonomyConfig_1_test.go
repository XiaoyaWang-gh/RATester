package hugolib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetTaxonomyConfig_1(t *testing.T) {
	cfg := contentMapConfig{
		lang: "en",
		taxonomyConfig: taxonomiesConfigValues{
			views: []viewName{
				{
					singular:      "category",
					plural:        "categories",
					pluralTreeKey: "categories",
				},
			},
		},
	}

	tests := []struct {
		name string
		cfg  contentMapConfig
		arg  string
		want viewName
	}{
		{
			name: "Test case 1",
			cfg:  cfg,
			arg:  "categories",
			want: viewName{
				singular:      "category",
				plural:        "categories",
				pluralTreeKey: "categories",
			},
		},
		{
			name: "Test case 2",
			cfg:  cfg,
			arg:  "tags",
			want: viewName{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.cfg.getTaxonomyConfig(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTaxonomyConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

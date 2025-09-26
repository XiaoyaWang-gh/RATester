package config

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
	"gotest.tools/assert"
)

func TestSetDefaultMergeStrategy_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		provider *defaultConfigProvider
		expected *defaultConfigProvider
	}{
		{
			name: "Test SetDefaultMergeStrategy with no existing merge strategy",
			provider: &defaultConfigProvider{
				root: maps.Params{
					"key1": "value1",
					"key2": "value2",
				},
			},
			expected: &defaultConfigProvider{
				root: maps.Params{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			name: "Test SetDefaultMergeStrategy with existing merge strategy",
			provider: &defaultConfigProvider{
				root: maps.Params{
					"key1":          "value1",
					"key2":          "value2",
					"mergeStrategy": "override",
				},
			},
			expected: &defaultConfigProvider{
				root: maps.Params{
					"key1":          "value1",
					"key2":          "value2",
					"mergeStrategy": "override",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			test.provider.SetDefaultMergeStrategy()
			assert.Equal(t, test.expected, test.provider)
		})
	}
}

package hugolib

import (
	"fmt"
	"testing"
)

func TestshouldRenderStandalonePage_1(t *testing.T) {
	type testCase struct {
		name     string
		context  siteRenderContext
		kind     string
		expected bool
	}

	testCases := []testCase{
		{
			name: "Test case 1",
			context: siteRenderContext{
				cfg: &BuildCfg{
					SkipRender: false,
				},
				languageIdx: 0,
				sitesOutIdx: 0,
				outIdx:      0,
				multihost:   false,
			},
			kind:     "KindSitemap",
			expected: true,
		},
		{
			name: "Test case 2",
			context: siteRenderContext{
				cfg: &BuildCfg{
					SkipRender: true,
				},
				languageIdx: 0,
				sitesOutIdx: 0,
				outIdx:      0,
				multihost:   false,
			},
			kind:     "KindStatus404",
			expected: false,
		},
		{
			name: "Test case 3",
			context: siteRenderContext{
				cfg: &BuildCfg{
					SkipRender: false,
				},
				languageIdx: 1,
				sitesOutIdx: 1,
				outIdx:      1,
				multihost:   true,
			},
			kind:     "KindSitemap",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.context.shouldRenderStandalonePage(tc.kind)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

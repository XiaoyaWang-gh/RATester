package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetSameSite_1(t *testing.T) {
	testCases := []struct {
		name     string
		samesite http.SameSite
		expected http.SameSite
	}{
		{
			name:     "TestSetSameSite_Case1",
			samesite: http.SameSiteDefaultMode,
			expected: http.SameSiteDefaultMode,
		},
		{
			name:     "TestSetSameSite_Case2",
			samesite: http.SameSiteLaxMode,
			expected: http.SameSiteLaxMode,
		},
		{
			name:     "TestSetSameSite_Case3",
			samesite: http.SameSiteStrictMode,
			expected: http.SameSiteStrictMode,
		},
		{
			name:     "TestSetSameSite_Case4",
			samesite: http.SameSiteNoneMode,
			expected: http.SameSiteNoneMode,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := &Context{}
			ctx.SetSameSite(tc.samesite)
			if ctx.sameSite != tc.expected {
				t.Errorf("Expected sameSite to be %v, but got %v", tc.expected, ctx.sameSite)
			}
		})
	}
}

package langs

import (
	"fmt"
	"testing"
)

func TestLanguageCode_1(t *testing.T) {
	type test struct {
		name     string
		lang     string
		expected string
	}

	tests := []test{
		{"English", "en", "en"},
		{"Norwegian", "no", "no"},
		{"Empty", "", "en"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &Language{
				Lang: tc.lang,
				LanguageConfig: LanguageConfig{
					LanguageCode: tc.expected,
				},
			}

			got := l.LanguageCode()

			if got != tc.expected {
				t.Errorf("Expected: %s, got: %s", tc.expected, got)
			}
		})
	}
}

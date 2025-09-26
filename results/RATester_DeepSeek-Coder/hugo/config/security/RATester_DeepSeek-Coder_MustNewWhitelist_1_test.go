package security

import (
	"fmt"
	"testing"
)

func TestMustNewWhitelist_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		patterns []string
		wantErr  bool
	}{
		{
			name:     "valid patterns",
			patterns: []string{"^a.*", "^b.*"},
			wantErr:  false,
		},
		{
			name:     "invalid pattern",
			patterns: []string{"[a-", "^b.*"},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt // NOTE: https://golang.org/doc/faq#closures_and_goroutines
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			got := func() (w Whitelist) {
				defer func() {
					if r := recover(); (r != nil) != tt.wantErr {
						t.Errorf("MustNewWhitelist() = %v, wantErr %v", r, tt.wantErr)
					}
				}()
				return MustNewWhitelist(tt.patterns...)
			}()

			if (got.patternsStrings == nil) != tt.wantErr {
				t.Errorf("MustNewWhitelist() = %v, wantErr %v", got, tt.wantErr)
			}
		})
	}
}

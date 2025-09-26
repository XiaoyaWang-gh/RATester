package httpcache

import (
	"fmt"
	"testing"
)

func TestCompilePredicate_1(t *testing.T) {
	tests := []struct {
		name    string
		gm      *GlobMatcher
		wantErr bool
	}{
		{
			name: "valid includes and excludes",
			gm: &GlobMatcher{
				Includes: []string{"*.txt", "*.go"},
				Excludes: []string{"*.go"},
			},
			wantErr: false,
		},
		{
			name: "invalid include pattern",
			gm: &GlobMatcher{
				Includes: []string{"[.txt", "*.go"},
				Excludes: []string{"*.go"},
			},
			wantErr: true,
		},
		{
			name: "invalid exclude pattern",
			gm: &GlobMatcher{
				Includes: []string{"*.txt", "*.go"},
				Excludes: []string{"[.go"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := tt.gm.CompilePredicate()
			if (err != nil) != tt.wantErr {
				t.Errorf("CompilePredicate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

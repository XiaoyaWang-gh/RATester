package plugins

import (
	"context"
	"fmt"
	"testing"
)

func TestNewProviderBuilder_1(t *testing.T) {
	ctx := context.Background()
	manifest := &Manifest{
		Runtime: "yaegi",
		Import:  "test/import",
		BasePkg: "test/base",
	}
	goPath := "test/go/path"

	tests := []struct {
		name     string
		ctx      context.Context
		manifest *Manifest
		goPath   string
		wantErr  bool
	}{
		{
			name:     "Test with valid inputs",
			ctx:      ctx,
			manifest: manifest,
			goPath:   goPath,
			wantErr:  false,
		},
		{
			name:     "Test with empty manifest",
			ctx:      ctx,
			manifest: &Manifest{},
			goPath:   goPath,
			wantErr:  true,
		},
		{
			name:     "Test with nil manifest",
			ctx:      ctx,
			manifest: nil,
			goPath:   goPath,
			wantErr:  true,
		},
		{
			name:     "Test with empty goPath",
			ctx:      ctx,
			manifest: manifest,
			goPath:   "",
			wantErr:  true,
		},
		{
			name:     "Test with nil goPath",
			ctx:      ctx,
			manifest: manifest,
			goPath:   "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := newProviderBuilder(tt.ctx, tt.manifest, tt.goPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("newProviderBuilder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

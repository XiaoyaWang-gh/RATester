package plugins

import (
	"context"
	"fmt"
	"testing"
)

func TestNewMiddlewareBuilder_1(t *testing.T) {
	ctx := context.Background()
	goPath := "/path/to/go"
	moduleName := "moduleName"
	settings := Settings{
		Envs:   []string{"env1=val1", "env2=val2"},
		Mounts: []string{"/path/to/mount1", "/path/to/mount2"},
	}

	tests := []struct {
		name      string
		manifest  *Manifest
		wantError bool
	}{
		{
			name: "WasmRuntime",
			manifest: &Manifest{
				Runtime: runtimeWasm,
			},
			wantError: false,
		},
		{
			name: "YaegiRuntime",
			manifest: &Manifest{
				Runtime: "",
			},
			wantError: false,
		},
		{
			name: "UnknownRuntime",
			manifest: &Manifest{
				Runtime: "unknown",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := newMiddlewareBuilder(ctx, goPath, tt.manifest, moduleName, settings)
			if (err != nil) != tt.wantError {
				t.Errorf("newMiddlewareBuilder() error = %v, wantError %v", err, tt.wantError)
				return
			}
		})
	}
}

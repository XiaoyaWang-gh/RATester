package plugins

import (
	"context"
	"fmt"
	"testing"

	"github.com/tetratelabs/wazero"
)

func TestInstantiateHost_1(t *testing.T) {
	tests := []struct {
		name     string
		runtime  wazero.Runtime
		mod      wazero.CompiledModule
		settings Settings
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := InstantiateHost(context.Background(), tt.runtime, tt.mod, tt.settings)
			if (err != nil) != tt.wantErr {
				t.Errorf("InstantiateHost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

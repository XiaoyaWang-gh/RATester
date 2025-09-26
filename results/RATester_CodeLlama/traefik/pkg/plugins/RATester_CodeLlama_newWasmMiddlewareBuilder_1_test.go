package plugins

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/tetratelabs/wazero"
)

func TestNewWasmMiddlewareBuilder_1(t *testing.T) {
	type args struct {
		goPath     string
		moduleName string
		wasmPath   string
		settings   Settings
	}
	tests := []struct {
		name    string
		args    args
		want    *wasmMiddlewareBuilder
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				goPath:     "test",
				moduleName: "test",
				wasmPath:   "test",
				settings: Settings{
					Envs: []string{"test"},
				},
			},
			want: &wasmMiddlewareBuilder{
				path:  "test",
				cache: wazero.NewCompilationCache(),
				settings: Settings{
					Envs: []string{"test"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := newWasmMiddlewareBuilder(tt.args.goPath, tt.args.moduleName, tt.args.wasmPath, tt.args.settings)
			if (err != nil) != tt.wantErr {
				t.Errorf("newWasmMiddlewareBuilder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newWasmMiddlewareBuilder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

package plugins

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReadManifest_1(t *testing.T) {
	type args struct {
		goPath     string
		moduleName string
	}
	tests := []struct {
		name    string
		args    args
		want    *Manifest
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				goPath:     "/path/to/go",
				moduleName: "module1",
			},
			want: &Manifest{
				DisplayName:   "Test Plugin",
				Type:          "wasm",
				Runtime:       "yaegi",
				WasmPath:      "plugin.wasm",
				Import:        "plugin",
				BasePkg:       "base",
				Compatibility: "1.0",
				Summary:       "This is a test plugin",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				goPath:     "/path/to/go",
				moduleName: "module2",
			},
			want:    nil,
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

			got, err := ReadManifest(tt.args.goPath, tt.args.moduleName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadManifest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadManifest() = %v, want %v", got, tt.want)
			}
		})
	}
}

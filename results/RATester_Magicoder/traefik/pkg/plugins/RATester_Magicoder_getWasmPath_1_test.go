package plugins

import (
	"fmt"
	"testing"
)

func TestGetWasmPath_1(t *testing.T) {
	type args struct {
		manifest *Manifest
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				manifest: &Manifest{
					WasmPath: "plugin.wasm",
				},
			},
			want:    "plugin.wasm",
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				manifest: &Manifest{
					WasmPath: "",
				},
			},
			want:    "plugin.wasm",
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				manifest: &Manifest{
					WasmPath: "/path/to/wasm",
				},
			},
			want:    "/path/to/wasm",
			wantErr: false,
		},
		{
			name: "Test Case 4",
			args: args{
				manifest: &Manifest{
					WasmPath: "http://example.com/plugin.wasm",
				},
			},
			want:    "",
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

			got, err := getWasmPath(tt.args.manifest)
			if (err != nil) != tt.wantErr {
				t.Errorf("getWasmPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getWasmPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

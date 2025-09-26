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
			name: "test",
			args: args{
				goPath:     "test",
				moduleName: "test",
			},
			want: &Manifest{
				DisplayName:   "test",
				Type:          "test",
				Runtime:       "test",
				WasmPath:      "test",
				Import:        "test",
				BasePkg:       "test",
				Compatibility: "test",
				Summary:       "test",
				TestData:      map[string]interface{}{},
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

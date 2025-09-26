package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestLoadConfigFromFile_1(t *testing.T) {
	type args struct {
		fs       afero.Fs
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]any
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				fs:       afero.NewMemMapFs(),
				filename: "test.yaml",
			},
			want: map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				fs:       afero.NewMemMapFs(),
				filename: "test.yaml",
			},
			want: map[string]any{
				"key1": "value1",
				"key2": "value2",
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

			got, err := loadConfigFromFile(tt.args.fs, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadConfigFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadConfigFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

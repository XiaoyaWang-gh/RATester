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
		// TODO: Add test cases.
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

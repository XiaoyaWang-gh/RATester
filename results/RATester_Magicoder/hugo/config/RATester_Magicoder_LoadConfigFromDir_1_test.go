package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestLoadConfigFromDir_1(t *testing.T) {
	type args struct {
		sourceFs    afero.Fs
		configDir   string
		environment string
	}
	tests := []struct {
		name    string
		args    args
		want    Provider
		want1   []string
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

			got, got1, err := LoadConfigFromDir(tt.args.sourceFs, tt.args.configDir, tt.args.environment)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfigFromDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfigFromDir() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LoadConfigFromDir() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

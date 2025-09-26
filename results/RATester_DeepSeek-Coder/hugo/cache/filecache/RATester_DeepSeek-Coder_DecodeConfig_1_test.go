package filecache

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/spf13/afero"
)

func TestDecodeConfig_1(t *testing.T) {
	type args struct {
		fs   afero.Fs
		bcfg config.BaseConfig
		m    map[string]any
	}
	tests := []struct {
		name    string
		args    args
		want    Configs
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

			got, err := DecodeConfig(tt.args.fs, tt.args.bcfg, tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

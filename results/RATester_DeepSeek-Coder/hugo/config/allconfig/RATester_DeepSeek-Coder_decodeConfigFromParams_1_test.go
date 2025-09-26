package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/config"
	"github.com/spf13/afero"
)

func TestDecodeConfigFromParams_1(t *testing.T) {
	type args struct {
		fs     afero.Fs
		logger loggers.Logger
		bcfg   config.BaseConfig
		p      config.Provider
		target *Config
		keys   []string
	}
	tests := []struct {
		name    string
		args    args
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

			if err := decodeConfigFromParams(tt.args.fs, tt.args.logger, tt.args.bcfg, tt.args.p, tt.args.target, tt.args.keys); (err != nil) != tt.wantErr {
				t.Errorf("decodeConfigFromParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

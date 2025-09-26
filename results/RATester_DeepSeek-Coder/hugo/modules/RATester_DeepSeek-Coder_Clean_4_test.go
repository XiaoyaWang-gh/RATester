package modules

import (
	"fmt"
	"testing"

	"github.com/gobwas/glob"
	"github.com/gohugoio/hugo/common/loggers"
	"github.com/spf13/afero"
)

func TestClean_4(t *testing.T) {
	type fields struct {
		fs                afero.Fs
		logger            loggers.Logger
		noVendor          glob.Glob
		ccfg              ClientConfig
		moduleConfig      Config
		environ           []string
		GoModulesFilename string
		goBinaryStatus    goBinaryStatus
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name    string
		fields  fields
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

			c := &Client{
				fs:                tt.fields.fs,
				logger:            tt.fields.logger,
				noVendor:          tt.fields.noVendor,
				ccfg:              tt.fields.ccfg,
				moduleConfig:      tt.fields.moduleConfig,
				environ:           tt.fields.environ,
				GoModulesFilename: tt.fields.GoModulesFilename,
				goBinaryStatus:    tt.fields.goBinaryStatus,
			}
			if err := c.Clean(tt.args.pattern); (err != nil) != tt.wantErr {
				t.Errorf("Client.Clean() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

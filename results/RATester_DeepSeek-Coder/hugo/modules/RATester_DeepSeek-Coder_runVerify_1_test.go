package modules

import (
	"fmt"
	"testing"

	"github.com/gobwas/glob"
	"github.com/gohugoio/hugo/common/loggers"
	"github.com/spf13/afero"
)

func TestRunVerify_1(t *testing.T) {
	t.Parallel()

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

	tests := []struct {
		name    string
		fields  fields
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
			if err := c.runVerify(); (err != nil) != tt.wantErr {
				t.Errorf("Client.runVerify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

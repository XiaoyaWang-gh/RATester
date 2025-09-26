package highlight

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
)

func TestApplyLegacyConfig_1(t *testing.T) {
	type args struct {
		cfg  config.Provider
		conf *Config
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

			if err := ApplyLegacyConfig(tt.args.cfg, tt.args.conf); (err != nil) != tt.wantErr {
				t.Errorf("ApplyLegacyConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

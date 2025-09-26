package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/modules"
)

func TestApplyConfigAliases_1(t *testing.T) {
	type fields struct {
		cfg        config.Provider
		BaseConfig config.BaseConfig
		ConfigSourceDescriptor
		ModulesConfig      modules.ModulesConfig
		ModulesConfigFiles []string
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

			l := configLoader{
				cfg:                    tt.fields.cfg,
				BaseConfig:             tt.fields.BaseConfig,
				ConfigSourceDescriptor: tt.fields.ConfigSourceDescriptor,
				ModulesConfig:          tt.fields.ModulesConfig,
				ModulesConfigFiles:     tt.fields.ModulesConfigFiles,
			}
			if err := l.applyConfigAliases(); (err != nil) != tt.wantErr {
				t.Errorf("applyConfigAliases() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

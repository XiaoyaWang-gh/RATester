package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/modules"
)

func TestDeleteMergeStrategies_1(t *testing.T) {
	type fields struct {
		cfg        config.Provider
		BaseConfig config.BaseConfig
		ConfigSourceDescriptor
		ModulesConfig      modules.ModulesConfig
		ModulesConfigFiles []string
	}
	tests := []struct {
		name   string
		fields fields
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
			l.deleteMergeStrategies()
		})
	}
}

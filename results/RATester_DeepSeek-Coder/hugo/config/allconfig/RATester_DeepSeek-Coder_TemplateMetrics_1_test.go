package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/langs"
)

func TestTemplateMetrics_1(t *testing.T) {
	type fields struct {
		config     *Config
		baseConfig config.BaseConfig
		m          *Configs
		language   *langs.Language
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
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

			c := ConfigLanguage{
				config:     tt.fields.config,
				baseConfig: tt.fields.baseConfig,
				m:          tt.fields.m,
				language:   tt.fields.language,
			}
			if got := c.TemplateMetrics(); got != tt.want {
				t.Errorf("ConfigLanguage.TemplateMetrics() = %v, want %v", got, tt.want)
			}
		})
	}
}

package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/langs"
)

func TestIsLangDisabled_1(t *testing.T) {
	type fields struct {
		config     *Config
		baseConfig config.BaseConfig
		m          *Configs
		language   *langs.Language
	}
	type args struct {
		lang string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
			if got := c.IsLangDisabled(tt.args.lang); got != tt.want {
				t.Errorf("IsLangDisabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

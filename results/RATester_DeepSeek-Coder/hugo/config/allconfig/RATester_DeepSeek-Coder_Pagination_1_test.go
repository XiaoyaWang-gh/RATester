package allconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/langs"
)

func TestPagination_1(t *testing.T) {
	type fields struct {
		config     *Config
		baseConfig config.BaseConfig
		m          *Configs
		language   *langs.Language
	}
	tests := []struct {
		name   string
		fields fields
		want   config.Pagination
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
			if got := c.Pagination(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pagination() = %v, want %v", got, tt.want)
			}
		})
	}
}

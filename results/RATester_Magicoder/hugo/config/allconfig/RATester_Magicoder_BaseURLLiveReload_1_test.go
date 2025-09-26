package allconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/urls"
)

func TestBaseURLLiveReload_1(t *testing.T) {
	type fields struct {
		config *Config
	}
	tests := []struct {
		name   string
		fields fields
		want   urls.BaseURL
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
				config: tt.fields.config,
			}
			if got := c.BaseURLLiveReload(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseURLLiveReload() = %v, want %v", got, tt.want)
			}
		})
	}
}

package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func Testinit_45(t *testing.T) {
	type fields struct {
		Params maps.Params
		Target PageMatcher
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

			p := &PageMatcherParamsConfig{
				Params: tt.fields.Params,
				Target: tt.fields.Target,
			}
			if err := p.init(); (err != nil) != tt.wantErr {
				t.Errorf("PageMatcherParamsConfig.init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

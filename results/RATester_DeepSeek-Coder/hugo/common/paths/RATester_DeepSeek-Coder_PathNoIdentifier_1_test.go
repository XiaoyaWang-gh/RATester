package paths

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types"
)

func TestPathNoIdentifier_1(t *testing.T) {
	type fields struct {
		s                     string
		posContainerLow       int
		posContainerHigh      int
		posSectionHigh        int
		component             string
		bundleType            PathType
		identifiers           []types.LowHigh[string]
		posIdentifierLanguage int
		disabled              bool
		trimLeadingSlash      bool
		unnormalized          *Path
	}
	tests := []struct {
		name   string
		fields fields
		want   string
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

			p := &Path{
				s:                     tt.fields.s,
				posContainerLow:       tt.fields.posContainerLow,
				posContainerHigh:      tt.fields.posContainerHigh,
				posSectionHigh:        tt.fields.posSectionHigh,
				component:             tt.fields.component,
				bundleType:            tt.fields.bundleType,
				identifiers:           tt.fields.identifiers,
				posIdentifierLanguage: tt.fields.posIdentifierLanguage,
				disabled:              tt.fields.disabled,
				trimLeadingSlash:      tt.fields.trimLeadingSlash,
				unnormalized:          tt.fields.unnormalized,
			}
			if got := p.PathNoIdentifier(); got != tt.want {
				t.Errorf("Path.PathNoIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

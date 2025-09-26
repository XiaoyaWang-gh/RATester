package page

import (
	"fmt"
	"testing"
)

func TestPathDirBase_1(t *testing.T) {
	type fields struct {
		els                []string
		d                  TargetPathDescriptor
		isUgly             bool
		baseNameSameAsType bool
		noSubResources     bool
		fullSuffix         string
		prefixLink         string
		prefixPath         string
		linkUpperOffset    int
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

			p := &pagePathBuilder{
				els:                tt.fields.els,
				d:                  tt.fields.d,
				isUgly:             tt.fields.isUgly,
				baseNameSameAsType: tt.fields.baseNameSameAsType,
				noSubResources:     tt.fields.noSubResources,
				fullSuffix:         tt.fields.fullSuffix,
				prefixLink:         tt.fields.prefixLink,
				prefixPath:         tt.fields.prefixPath,
				linkUpperOffset:    tt.fields.linkUpperOffset,
			}
			if got := p.PathDirBase(); got != tt.want {
				t.Errorf("pagePathBuilder.PathDirBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

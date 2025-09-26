package legacy

import (
	"fmt"
	"testing"

	"gopkg.in/ini.v1"
)

func TestRenderRangeProxyTemplates_1(t *testing.T) {
	type args struct {
		f       *ini.File
		section *ini.Section
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

			if err := renderRangeProxyTemplates(tt.args.f, tt.args.section); (err != nil) != tt.wantErr {
				t.Errorf("renderRangeProxyTemplates() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

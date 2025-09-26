package modules

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
)

func TestApplyProjectConfigDefaults_1(t *testing.T) {
	type args struct {
		mod  Module
		cfgs []config.AllProvider
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

			if err := ApplyProjectConfigDefaults(tt.args.mod, tt.args.cfgs...); (err != nil) != tt.wantErr {
				t.Errorf("ApplyProjectConfigDefaults() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

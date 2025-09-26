package file

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestapplyConfiguration_1(t *testing.T) {
	type fields struct {
		Directory                 string
		Watch                     bool
		Filename                  string
		DebugLogGeneratedTemplate bool
	}
	type args struct {
		configurationChan chan<- dynamic.Message
	}
	tests := []struct {
		name    string
		fields  fields
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

			p := &Provider{
				Directory:                 tt.fields.Directory,
				Watch:                     tt.fields.Watch,
				Filename:                  tt.fields.Filename,
				DebugLogGeneratedTemplate: tt.fields.DebugLogGeneratedTemplate,
			}
			if err := p.applyConfiguration(tt.args.configurationChan); (err != nil) != tt.wantErr {
				t.Errorf("Provider.applyConfiguration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

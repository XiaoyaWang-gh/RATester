package file

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestbuildConfiguration_5(t *testing.T) {
	tests := []struct {
		name    string
		fields  Provider
		want    *dynamic.Configuration
		wantErr bool
	}{
		{
			name: "Test case 1",
			fields: Provider{
				Directory: "/path/to/directory",
			},
			want:    &dynamic.Configuration{},
			wantErr: false,
		},
		{
			name: "Test case 2",
			fields: Provider{
				Filename: "/path/to/file.yml",
			},
			want:    &dynamic.Configuration{},
			wantErr: false,
		},
		{
			name: "Test case 3",
			fields: Provider{
				Directory: "",
				Filename:  "",
			},
			want:    nil,
			wantErr: true,
		},
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
			got, err := p.buildConfiguration()
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.buildConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.buildConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

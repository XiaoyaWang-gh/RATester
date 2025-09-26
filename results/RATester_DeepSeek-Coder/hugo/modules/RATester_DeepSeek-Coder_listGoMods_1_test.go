package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestListGoMods_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		client  *Client
		want    goModules
		wantErr bool
	}{
		{
			name: "success",
			client: &Client{
				GoModulesFilename: "go.mod",
				moduleConfig: Config{
					Imports: []Import{{
						Path: "github.com/gohugoio/hugo",
					}},
				},
			},
			want:    goModules{{Path: "github.com/gohugoio/hugo"}},
			wantErr: false,
		},
		{
			name: "no go.mod",
			client: &Client{
				GoModulesFilename: "",
				moduleConfig: Config{
					Imports: []Import{{
						Path: "github.com/gohugoio/hugo",
					}},
				},
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "no module import",
			client: &Client{
				GoModulesFilename: "go.mod",
				moduleConfig: Config{
					Imports: []Import{},
				},
			},
			want:    nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.client.listGoMods()
			if (err != nil) != tt.wantErr {
				t.Errorf("listGoMods() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listGoMods() = %v, want %v", got, tt.want)
			}
		})
	}
}

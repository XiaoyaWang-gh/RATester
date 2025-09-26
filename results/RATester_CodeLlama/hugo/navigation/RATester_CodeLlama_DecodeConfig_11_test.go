package navigation

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
)

func TestDecodeConfig_11(t *testing.T) {
	type args struct {
		in any
	}
	tests := []struct {
		name    string
		args    args
		want    *config.ConfigNamespace[map[string]MenuConfig, Menus]
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				in: nil,
			},
			want: &config.ConfigNamespace[map[string]MenuConfig, Menus]{
				SourceStructure: nil,
				SourceHash:      "",
				Config:          Menus{},
			},
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

			got, err := DecodeConfig(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

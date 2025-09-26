package output

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/media"
)

func TestDecodeConfig_12(t *testing.T) {
	type args struct {
		mediaTypes media.Types
		in         any
	}
	tests := []struct {
		name    string
		args    args
		want    *config.ConfigNamespace[map[string]OutputFormatConfig, Formats]
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

			got, err := DecodeConfig(tt.args.mediaTypes, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

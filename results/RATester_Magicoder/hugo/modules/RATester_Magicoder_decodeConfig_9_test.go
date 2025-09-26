package modules

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
)

func TestdecodeConfig_9(t *testing.T) {
	type args struct {
		cfg              config.Provider
		pathReplacements map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    Config
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

			got, err := decodeConfig(tt.args.cfg, tt.args.pathReplacements)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

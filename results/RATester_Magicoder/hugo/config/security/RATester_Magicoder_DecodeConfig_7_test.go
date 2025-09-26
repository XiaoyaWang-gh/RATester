package security

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
)

func TestDecodeConfig_7(t *testing.T) {
	type args struct {
		cfg config.Provider
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

			got, err := DecodeConfig(tt.args.cfg)
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

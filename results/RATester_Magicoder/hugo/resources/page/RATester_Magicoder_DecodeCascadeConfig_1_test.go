package page

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/common/maps"
	"github.com/gohugoio/hugo/config"
)

func TestDecodeCascadeConfig_1(t *testing.T) {
	type args struct {
		logger loggers.Logger
		in     any
	}
	tests := []struct {
		name    string
		args    args
		want    *config.ConfigNamespace[[]PageMatcherParamsConfig, map[PageMatcher]maps.Params]
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

			got, err := DecodeCascadeConfig(tt.args.logger, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeCascadeConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeCascadeConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

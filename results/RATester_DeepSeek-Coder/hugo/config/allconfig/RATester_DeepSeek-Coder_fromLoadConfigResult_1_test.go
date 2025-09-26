package allconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/config"
	"github.com/spf13/afero"
)

func TestFromLoadConfigResult_1(t *testing.T) {
	type args struct {
		fs     afero.Fs
		logger loggers.Logger
		res    config.LoadConfigResult
	}
	tests := []struct {
		name    string
		args    args
		want    *Configs
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

			got, err := fromLoadConfigResult(tt.args.fs, tt.args.logger, tt.args.res)
			if (err != nil) != tt.wantErr {
				t.Errorf("fromLoadConfigResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fromLoadConfigResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

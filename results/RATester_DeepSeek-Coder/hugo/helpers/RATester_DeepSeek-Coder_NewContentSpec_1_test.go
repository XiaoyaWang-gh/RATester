package helpers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/hexec"
	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/config"
	"github.com/spf13/afero"
)

func TestNewContentSpec_1(t *testing.T) {
	type args struct {
		cfg       config.AllProvider
		logger    loggers.Logger
		contentFs afero.Fs
		ex        *hexec.Exec
	}
	tests := []struct {
		name    string
		args    args
		want    *ContentSpec
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

			got, err := NewContentSpec(tt.args.cfg, tt.args.logger, tt.args.contentFs, tt.args.ex)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewContentSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContentSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

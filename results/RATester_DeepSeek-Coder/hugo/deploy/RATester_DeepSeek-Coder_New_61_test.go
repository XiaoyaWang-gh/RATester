package deploy

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/config"
	"github.com/spf13/afero"
)

func TestNew_61(t *testing.T) {
	type args struct {
		cfg     config.AllProvider
		logger  loggers.Logger
		localFs afero.Fs
	}
	tests := []struct {
		name    string
		args    args
		want    *Deployer
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

			got, err := New(tt.args.cfg, tt.args.logger, tt.args.localFs)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

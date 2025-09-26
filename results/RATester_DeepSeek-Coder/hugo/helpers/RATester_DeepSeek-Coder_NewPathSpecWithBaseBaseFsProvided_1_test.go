package helpers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugolib/filesystems"
)

func TestNewPathSpecWithBaseBaseFsProvided_1(t *testing.T) {
	t.Parallel()

	type args struct {
		fs         *hugofs.Fs
		cfg        config.AllProvider
		logger     loggers.Logger
		baseBaseFs *filesystems.BaseFs
	}

	tests := []struct {
		name    string
		args    args
		want    *PathSpec
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

			got, err := NewPathSpecWithBaseBaseFsProvided(tt.args.fs, tt.args.cfg, tt.args.logger, tt.args.baseBaseFs)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPathSpecWithBaseBaseFsProvided() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPathSpecWithBaseBaseFsProvided() = %v, want %v", got, tt.want)
			}
		})
	}
}

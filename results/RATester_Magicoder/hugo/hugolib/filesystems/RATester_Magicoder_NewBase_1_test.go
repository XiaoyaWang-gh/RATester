package filesystems

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/hugolib/paths"
)

func TestNewBase_1(t *testing.T) {
	type args struct {
		p       *paths.Paths
		logger  loggers.Logger
		options []func(*BaseFs) error
	}
	tests := []struct {
		name    string
		args    args
		want    *BaseFs
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

			got, err := NewBase(tt.args.p, tt.args.logger, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

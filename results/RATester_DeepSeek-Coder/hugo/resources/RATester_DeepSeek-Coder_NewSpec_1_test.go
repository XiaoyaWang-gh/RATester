package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/cache/dynacache"
	"github.com/gohugoio/hugo/cache/filecache"
	"github.com/gohugoio/hugo/common/herrors"
	"github.com/gohugoio/hugo/common/hexec"
	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/common/types"
	"github.com/gohugoio/hugo/helpers"
	"github.com/gohugoio/hugo/identity"
)

func TestNewSpec_1(t *testing.T) {
	type args struct {
		s            *helpers.PathSpec
		common       *SpecCommon
		fileCaches   filecache.Caches
		memCache     *dynacache.Cache
		incr         identity.Incrementer
		logger       loggers.Logger
		errorHandler herrors.ErrorSender
		execHelper   *hexec.Exec
		buildClosers types.CloseAdder
		rebuilder    identity.SignalRebuilder
	}
	tests := []struct {
		name    string
		args    args
		want    *Spec
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

			got, err := NewSpec(tt.args.s, tt.args.common, tt.args.fileCaches, tt.args.memCache, tt.args.incr, tt.args.logger, tt.args.errorHandler, tt.args.execHelper, tt.args.buildClosers, tt.args.rebuilder)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSpec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

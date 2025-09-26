package page

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/common/maps"
)

func TestDecodeCascade_1(t *testing.T) {
	type args struct {
		logger loggers.Logger
		in     any
	}
	tests := []struct {
		name    string
		args    args
		want    map[PageMatcher]maps.Params
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

			got, err := DecodeCascade(tt.args.logger, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeCascade() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeCascade() = %v, want %v", got, tt.want)
			}
		})
	}
}

package minifiers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/media"
	"github.com/gohugoio/hugo/output"
)

func TestNew_23(t *testing.T) {
	type args struct {
		mediaTypes    media.Types
		outputFormats output.Formats
		cfg           config.AllProvider
	}
	tests := []struct {
		name    string
		args    args
		want    Client
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

			got, err := New(tt.args.mediaTypes, tt.args.outputFormats, tt.args.cfg)
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

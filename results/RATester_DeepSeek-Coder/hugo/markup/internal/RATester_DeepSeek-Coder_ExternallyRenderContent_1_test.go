package internal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestExternallyRenderContent_1(t *testing.T) {
	type args struct {
		cfg        converter.ProviderConfig
		ctx        converter.DocumentContext
		content    []byte
		binaryName string
		args       []string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
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

			got, err := ExternallyRenderContent(tt.args.cfg, tt.args.ctx, tt.args.content, tt.args.binaryName, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExternallyRenderContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExternallyRenderContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

package internal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/tableofcontents"
)

func TestextractTOC_2(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		a       *AsciidocConverter
		args    args
		want    []byte
		want1   *tableofcontents.Fragments
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

			got, got1, err := tt.a.extractTOC(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("AsciidocConverter.extractTOC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AsciidocConverter.extractTOC() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("AsciidocConverter.extractTOC() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

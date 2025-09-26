package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/converter"
)

func TestnewContentConverter_1(t *testing.T) {
	type args struct {
		ps     *pageState
		markup string
	}
	tests := []struct {
		name    string
		p       *pageMeta
		args    args
		want    converter.Converter
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

			got, err := tt.p.newContentConverter(tt.args.ps, tt.args.markup)
			if (err != nil) != tt.wantErr {
				t.Errorf("newContentConverter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newContentConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}

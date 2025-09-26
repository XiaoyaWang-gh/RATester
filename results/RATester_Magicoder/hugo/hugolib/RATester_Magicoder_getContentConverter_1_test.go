package hugolib

import (
	"fmt"
	"reflect"
	"sync/atomic"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
	"github.com/gohugoio/hugo/markup/converter"
)

func TestgetContentConverter_1(t *testing.T) {
	type fields struct {
		po                     *pageOutput
		otherOutputs           *maps.Cache[uint64, *pageContentOutput]
		contentRenderedVersion uint32
		contentRendered        atomic.Bool
		renderHooks            *renderHooks
	}
	tests := []struct {
		name    string
		fields  fields
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

			pco := &pageContentOutput{
				po:                     tt.fields.po,
				otherOutputs:           tt.fields.otherOutputs,
				contentRenderedVersion: tt.fields.contentRenderedVersion,
				contentRendered:        tt.fields.contentRendered,
				renderHooks:            tt.fields.renderHooks,
			}
			got, err := pco.getContentConverter()
			if (err != nil) != tt.wantErr {
				t.Errorf("getContentConverter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getContentConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}

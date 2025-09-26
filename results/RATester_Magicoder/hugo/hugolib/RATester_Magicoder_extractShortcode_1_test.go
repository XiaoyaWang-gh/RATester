package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/parser/pageparser"
)

func TestextractShortcode_1(t *testing.T) {
	tests := []struct {
		name    string
		handler *shortcodeHandler
		ordinal int
		level   int
		source  []byte
		pt      *pageparser.Iterator
		want    *shortcode
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

			s := &shortcodeHandler{}
			got, err := s.extractShortcode(tt.ordinal, tt.level, tt.source, tt.pt)
			if (err != nil) != tt.wantErr {
				t.Errorf("shortcodeHandler.extractShortcode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortcodeHandler.extractShortcode() = %v, want %v", got, tt.want)
			}
		})
	}
}

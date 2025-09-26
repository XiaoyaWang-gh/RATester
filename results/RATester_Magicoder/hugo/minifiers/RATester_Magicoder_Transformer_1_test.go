package minifiers

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/media"
	"github.com/gohugoio/hugo/transform"
)

func TestTransformer_1(t *testing.T) {
	type args struct {
		mediatype media.Type
	}
	tests := []struct {
		name string
		m    Client
		args args
		want transform.Transformer
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

			if got := tt.m.Transformer(tt.args.mediatype); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transformer() = %v, want %v", got, tt.want)
			}
		})
	}
}

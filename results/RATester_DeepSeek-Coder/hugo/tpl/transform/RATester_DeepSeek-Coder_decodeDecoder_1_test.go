package transform

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/parser/metadecoders"
)

func TestDecodeDecoder_1(t *testing.T) {
	type args struct {
		m map[string]any
	}
	tests := []struct {
		name    string
		args    args
		want    metadecoders.Decoder
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

			got, err := decodeDecoder(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeDecoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

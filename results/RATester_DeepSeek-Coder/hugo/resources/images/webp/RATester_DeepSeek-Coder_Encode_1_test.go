package webp

import (
	"fmt"
	"image"
	"io"
	"testing"

	"github.com/bep/gowebp/libwebp/webpoptions"
)

func TestEncode_1(t *testing.T) {
	type args struct {
		w io.Writer
		m image.Image
		o webpoptions.EncodingOptions
	}
	tests := []struct {
		name    string
		args    args
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

			if err := Encode(tt.args.w, tt.args.m, tt.args.o); (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

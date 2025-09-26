package exif

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/bep/imagemeta"
)

func TestDecode_2(t *testing.T) {
	type args struct {
		filename string
		format   imagemeta.ImageFormat
		r        io.Reader
	}
	tests := []struct {
		name    string
		d       *Decoder
		args    args
		wantEx  *ExifInfo
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

			gotEx, err := tt.d.Decode(tt.args.filename, tt.args.format, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEx, tt.wantEx) {
				t.Errorf("Decoder.Decode() = %v, want %v", gotEx, tt.wantEx)
			}
		})
	}
}

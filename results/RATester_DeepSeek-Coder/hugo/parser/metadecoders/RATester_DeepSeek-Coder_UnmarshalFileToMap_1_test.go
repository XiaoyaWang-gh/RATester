package metadecoders

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestUnmarshalFileToMap_1(t *testing.T) {
	type args struct {
		fs       afero.Fs
		filename string
	}
	tests := []struct {
		name    string
		d       Decoder
		args    args
		want    map[string]any
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

			got, err := tt.d.UnmarshalFileToMap(tt.args.fs, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decoder.UnmarshalFileToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decoder.UnmarshalFileToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/parser/metadecoders"
)

func TestReadConfig_1(t *testing.T) {
	type args struct {
		format metadecoders.Format
		data   []byte
	}
	tests := []struct {
		name    string
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

			got, err := readConfig(tt.args.format, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("readConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

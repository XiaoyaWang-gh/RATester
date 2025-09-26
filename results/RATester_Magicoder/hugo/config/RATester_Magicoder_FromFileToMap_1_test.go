package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestFromFileToMap_1(t *testing.T) {
	type args struct {
		fs       afero.Fs
		filename string
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

			got, err := FromFileToMap(tt.args.fs, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromFileToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromFileToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

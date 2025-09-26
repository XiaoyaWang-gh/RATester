package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/source"
	"github.com/spf13/afero"
)

func TestOpen_4(t *testing.T) {
	type fields struct {
		File *source.File
	}
	tests := []struct {
		name    string
		fields  fields
		want    afero.File
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

			fi := &fileInfo{
				File: tt.fields.File,
			}
			got, err := fi.Open()
			if (err != nil) != tt.wantErr {
				t.Errorf("fileInfo.Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fileInfo.Open() = %v, want %v", got, tt.want)
			}
		})
	}
}

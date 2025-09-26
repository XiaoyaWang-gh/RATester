package echo

import (
	"fmt"
	"io/fs"
	"reflect"
	"testing"
)

func TestOpen_1(t *testing.T) {
	type fields struct {
		fs     fs.FS
		prefix string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    fs.File
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

			fs := defaultFS{
				fs:     tt.fields.fs,
				prefix: tt.fields.prefix,
			}
			got, err := fs.Open(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultFS.Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defaultFS.Open() = %v, want %v", got, tt.want)
			}
		})
	}
}

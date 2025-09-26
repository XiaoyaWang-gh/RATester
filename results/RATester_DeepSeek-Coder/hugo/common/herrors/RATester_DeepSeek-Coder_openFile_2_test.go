package herrors

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestOpenFile_2(t *testing.T) {
	type args struct {
		filename string
		fs       afero.Fs
	}
	tests := []struct {
		name    string
		args    args
		want    afero.File
		want1   string
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

			got, got1, err := openFile(tt.args.filename, tt.args.fs)
			if (err != nil) != tt.wantErr {
				t.Errorf("openFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("openFile() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("openFile() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

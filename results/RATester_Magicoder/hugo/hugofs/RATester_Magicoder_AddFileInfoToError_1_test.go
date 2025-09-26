package hugofs

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestAddFileInfoToError_1(t *testing.T) {
	type args struct {
		err error
		fi  FileMetaInfo
		fs  afero.Fs
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

			if err := AddFileInfoToError(tt.args.err, tt.args.fi, tt.args.fs); (err != nil) != tt.wantErr {
				t.Errorf("AddFileInfoToError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

package filesystems

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestStatResource_1(t *testing.T) {
	type args struct {
		lang     string
		filename string
	}
	tests := []struct {
		name    string
		s       SourceFilesystems
		args    args
		wantFi  os.FileInfo
		wantFs  afero.Fs
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

			gotFi, gotFs, err := tt.s.StatResource(tt.args.lang, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("StatResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFi, tt.wantFi) {
				t.Errorf("StatResource() gotFi = %v, want %v", gotFi, tt.wantFi)
			}
			if !reflect.DeepEqual(gotFs, tt.wantFs) {
				t.Errorf("StatResource() gotFs = %v, want %v", gotFs, tt.wantFs)
			}
		})
	}
}

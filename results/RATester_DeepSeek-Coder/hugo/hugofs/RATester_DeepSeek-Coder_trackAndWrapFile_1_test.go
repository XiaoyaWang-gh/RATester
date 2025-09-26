package hugofs

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/spf13/afero"
)

func TestTrackAndWrapFile_1(t *testing.T) {
	type fields struct {
		Fs        afero.Fs
		mu        sync.Mutex
		openFiles map[string]int
	}
	type args struct {
		f afero.File
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   afero.File
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

			fs := &OpenFilesFs{
				Fs:        tt.fields.Fs,
				mu:        tt.fields.mu,
				openFiles: tt.fields.openFiles,
			}
			if got := fs.trackAndWrapFile(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenFilesFs.trackAndWrapFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

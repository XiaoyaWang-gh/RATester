package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestnewUnionFile_1(t *testing.T) {
	type args struct {
		fis []FileMetaInfo
	}
	tests := []struct {
		name    string
		args    args
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

			fs := &RootMappingFs{}
			got, err := fs.newUnionFile(tt.args.fis...)
			if (err != nil) != tt.wantErr {
				t.Errorf("newUnionFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newUnionFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

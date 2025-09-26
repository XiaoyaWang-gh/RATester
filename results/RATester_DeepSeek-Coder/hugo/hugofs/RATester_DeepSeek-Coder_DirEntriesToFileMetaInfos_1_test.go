package hugofs

import (
	"fmt"
	"io/fs"
	"reflect"
	"testing"
)

func TestDirEntriesToFileMetaInfos_1(t *testing.T) {
	type args struct {
		fis []fs.DirEntry
	}
	tests := []struct {
		name string
		args args
		want []FileMetaInfo
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

			if got := DirEntriesToFileMetaInfos(tt.args.fis); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DirEntriesToFileMetaInfos() = %v, want %v", got, tt.want)
			}
		})
	}
}

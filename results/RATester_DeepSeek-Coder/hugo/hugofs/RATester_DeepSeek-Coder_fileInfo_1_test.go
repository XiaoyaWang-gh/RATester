package hugofs

import (
	"fmt"
	"io/fs"
	"reflect"
	"sync"
	"testing"
)

func TestFileInfo_1(t *testing.T) {
	type fields struct {
		DirEntry fs.DirEntry
		m        *FileMeta
		fi       fs.FileInfo
		fiInit   sync.Once
	}
	tests := []struct {
		name   string
		fields fields
		want   fs.FileInfo
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

			fi := &dirEntryMeta{
				DirEntry: tt.fields.DirEntry,
				m:        tt.fields.m,
				fi:       tt.fields.fi,
				fiInit:   tt.fields.fiInit,
			}
			if got := fi.fileInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dirEntryMeta.fileInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

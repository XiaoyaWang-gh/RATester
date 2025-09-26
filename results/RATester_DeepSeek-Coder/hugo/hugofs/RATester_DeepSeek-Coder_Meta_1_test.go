package hugofs

import (
	"fmt"
	"io/fs"
	"reflect"
	"testing"
)

func TestMeta_1(t *testing.T) {
	type fields struct {
		DirEntry fs.DirEntry
		m        *FileMeta
	}
	tests := []struct {
		name   string
		fields fields
		want   *FileMeta
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
			}
			if got := fi.Meta(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dirEntryMeta.Meta() = %v, want %v", got, tt.want)
			}
		})
	}
}

package source

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TestFileInfo_3(t *testing.T) {
	type fields struct {
		fim      hugofs.FileMetaInfo
		uniqueID string
		lazyInit sync.Once
	}
	tests := []struct {
		name   string
		fields fields
		want   hugofs.FileMetaInfo
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

			fi := &File{
				fim:      tt.fields.fim,
				uniqueID: tt.fields.uniqueID,
				lazyInit: tt.fields.lazyInit,
			}
			if got := fi.FileInfo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File.FileInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

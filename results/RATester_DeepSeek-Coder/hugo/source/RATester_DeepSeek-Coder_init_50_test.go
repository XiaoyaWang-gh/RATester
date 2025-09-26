package source

import (
	"fmt"
	"sync"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TestInit_50(t *testing.T) {
	type fields struct {
		fim      hugofs.FileMetaInfo
		lazyInit sync.Once
	}
	tests := []struct {
		name   string
		fields fields
		want   string
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
				lazyInit: tt.fields.lazyInit,
			}
			fi.init()
			if fi.uniqueID != tt.want {
				t.Errorf("File.init() = %v, want %v", fi.uniqueID, tt.want)
			}
		})
	}
}

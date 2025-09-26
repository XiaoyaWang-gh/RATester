package source

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/gohugoio/hugo/common/hugio"
	"github.com/gohugoio/hugo/hugofs"
)

func TestOpen_5(t *testing.T) {
	type fields struct {
		fim      hugofs.FileMetaInfo
		uniqueID string
		lazyInit sync.Once
	}
	tests := []struct {
		name    string
		fields  fields
		want    hugio.ReadSeekCloser
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

			fi := &File{
				fim:      tt.fields.fim,
				uniqueID: tt.fields.uniqueID,
				lazyInit: tt.fields.lazyInit,
			}
			got, err := fi.Open()
			if (err != nil) != tt.wantErr {
				t.Errorf("File.Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File.Open() = %v, want %v", got, tt.want)
			}
		})
	}
}

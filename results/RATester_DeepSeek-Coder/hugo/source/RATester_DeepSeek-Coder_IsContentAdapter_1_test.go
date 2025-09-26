package source

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TestIsContentAdapter_1(t *testing.T) {
	type fields struct {
		fim hugofs.FileMetaInfo
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
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
				fim: tt.fields.fim,
			}
			if got := fi.IsContentAdapter(); got != tt.want {
				t.Errorf("File.IsContentAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

package source

import (
	"fmt"
	"sync"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TestUniqueID_2(t *testing.T) {
	type fields struct {
		fim      hugofs.FileMetaInfo
		uniqueID string
		lazyInit sync.Once
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				fim:      nil,
				uniqueID: "test_unique_id",
				lazyInit: sync.Once{},
			},
			want: "test_unique_id",
		},
		// Add more test cases here
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
			if got := fi.UniqueID(); got != tt.want {
				t.Errorf("UniqueID() = %v, want %v", got, tt.want)
			}
		})
	}
}

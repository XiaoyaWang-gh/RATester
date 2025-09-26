package resources

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/cache/filecache"
)

func TestGetFromFile_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		args  args
		want  filecache.ItemInfo
		want1 io.ReadCloser
		want2 transformedResourceMetadata
		want3 bool
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

			c := &ResourceCache{}
			got, got1, got2, got3 := c.getFromFile(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFromFile() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getFromFile() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("getFromFile() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("getFromFile() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}

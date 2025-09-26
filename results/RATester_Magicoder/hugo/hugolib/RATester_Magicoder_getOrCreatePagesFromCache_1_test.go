package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/cache/dynacache"
	"github.com/gohugoio/hugo/resources/page"
)

func TestgetOrCreatePagesFromCache_1(t *testing.T) {
	type args struct {
		cache  *dynacache.Partition[string, page.Pages]
		key    string
		create func(string) (page.Pages, error)
	}
	tests := []struct {
		name    string
		args    args
		want    page.Pages
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

			m := &pageMap{}
			got, err := m.getOrCreatePagesFromCache(tt.args.cache, tt.args.key, tt.args.create)
			if (err != nil) != tt.wantErr {
				t.Errorf("getOrCreatePagesFromCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOrCreatePagesFromCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

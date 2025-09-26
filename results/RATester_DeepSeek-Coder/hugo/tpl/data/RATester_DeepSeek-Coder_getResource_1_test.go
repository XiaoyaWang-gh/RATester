package data

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gohugoio/hugo/cache/filecache"
)

func TestGetResource_1(t *testing.T) {
	type args struct {
		cache     *filecache.Cache
		unmarshal func(b []byte) (bool, error)
		req       *http.Request
	}
	tests := []struct {
		name    string
		args    args
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

			ns := &Namespace{}
			if err := ns.getResource(tt.args.cache, tt.args.unmarshal, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Namespace.getResource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

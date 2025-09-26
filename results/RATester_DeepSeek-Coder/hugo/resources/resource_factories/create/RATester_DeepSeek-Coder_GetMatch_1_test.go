package create

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/bep/logg"
	"github.com/gohugoio/hugo/cache/filecache"
	hhttpcache "github.com/gohugoio/hugo/cache/httpcache"
	"github.com/gohugoio/hugo/common/hcontext"
	"github.com/gohugoio/hugo/common/tasks"
	"github.com/gohugoio/hugo/resources"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestGetMatch_1(t *testing.T) {
	type fields struct {
		rs                    *resources.Spec
		httpClient            *http.Client
		httpCacheConfig       hhttpcache.ConfigCompiled
		cacheGetResource      *filecache.Cache
		resourceIDDispatcher  hcontext.ContextDispatcher[string]
		remoteResourceChecker *tasks.RunEvery
		remoteResourceLogger  logg.LevelLogger
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    resource.Resource
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

			c := &Client{
				rs:                    tt.fields.rs,
				httpClient:            tt.fields.httpClient,
				httpCacheConfig:       tt.fields.httpCacheConfig,
				cacheGetResource:      tt.fields.cacheGetResource,
				resourceIDDispatcher:  tt.fields.resourceIDDispatcher,
				remoteResourceChecker: tt.fields.remoteResourceChecker,
				remoteResourceLogger:  tt.fields.remoteResourceLogger,
			}
			got, err := c.GetMatch(tt.args.pattern)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

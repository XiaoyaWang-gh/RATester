package echo

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"sync"
	"testing"
)

func TestLogger_2(t *testing.T) {
	type fields struct {
		logger   Logger
		request  *http.Request
		response *Response
		query    url.Values
		echo     *Echo
		store    Map
		lock     sync.RWMutex
		handler  HandlerFunc
		path     string
		pvalues  []string
		pnames   []string
	}
	tests := []struct {
		name   string
		fields fields
		want   Logger
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

			c := &context{
				logger:   tt.fields.logger,
				request:  tt.fields.request,
				response: tt.fields.response,
				query:    tt.fields.query,
				echo:     tt.fields.echo,
				store:    tt.fields.store,
				lock:     tt.fields.lock,
				handler:  tt.fields.handler,
				path:     tt.fields.path,
				pvalues:  tt.fields.pvalues,
				pnames:   tt.fields.pnames,
			}
			if got := c.Logger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("context.Logger() = %v, want %v", got, tt.want)
			}
		})
	}
}

package client

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestSearchCookieByKeyAndPath_1(t *testing.T) {
	type args struct {
		key     []byte
		path    []byte
		cookies []*fasthttp.Cookie
	}
	tests := []struct {
		name string
		args args
		want *fasthttp.Cookie
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

			if got := searchCookieByKeyAndPath(tt.args.key, tt.args.path, tt.args.cookies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchCookieByKeyAndPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

package urlrewrite

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestNewURLRewrite_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		next http.Handler
		conf dynamic.URLRewrite
		name string
	}
	tests := []struct {
		name string
		args args
		want http.Handler
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

			if got := NewURLRewrite(tt.args.ctx, tt.args.next, tt.args.conf, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewURLRewrite() = %v, want %v", got, tt.want)
			}
		})
	}
}

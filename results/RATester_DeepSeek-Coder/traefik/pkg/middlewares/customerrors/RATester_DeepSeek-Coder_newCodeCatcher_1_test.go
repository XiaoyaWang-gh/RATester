package customerrors

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestNewCodeCatcher_1(t *testing.T) {
	type args struct {
		rw             http.ResponseWriter
		httpCodeRanges types.HTTPCodeRanges
	}
	tests := []struct {
		name string
		args args
		want *codeCatcher
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

			if got := newCodeCatcher(tt.args.rw, tt.args.httpCodeRanges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCodeCatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}

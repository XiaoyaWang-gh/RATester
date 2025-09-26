package gin

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestWrapH_1(t *testing.T) {
	type args struct {
		h http.Handler
	}
	tests := []struct {
		name string
		args args
		want HandlerFunc
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

			if got := WrapH(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapH() = %v, want %v", got, tt.want)
			}
		})
	}
}

package i18n

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/source"
)

func TestErrWithFileContext_1(t *testing.T) {
	type args struct {
		inerr error
		r     *source.File
	}
	tests := []struct {
		name string
		args args
		want error
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

			if got := errWithFileContext(tt.args.inerr, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errWithFileContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

package bundler

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/hugio"
)

func TestnewMultiReadSeekCloser_1(t *testing.T) {
	type args struct {
		sources []hugio.ReadSeekCloser
	}
	tests := []struct {
		name string
		args args
		want *multiReadSeekCloser
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

			if got := newMultiReadSeekCloser(tt.args.sources...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMultiReadSeekCloser() = %v, want %v", got, tt.want)
			}
		})
	}
}

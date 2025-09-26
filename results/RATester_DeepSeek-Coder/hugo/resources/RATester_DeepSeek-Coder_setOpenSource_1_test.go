package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/hugio"
)

func TestSetOpenSource_1(t *testing.T) {
	type args struct {
		openSource hugio.OpenReadSeekCloser
	}
	tests := []struct {
		name string
		l    *genericResource
		args args
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

			l := &genericResource{}
			l.setOpenSource(tt.args.openSource)
		})
	}
}

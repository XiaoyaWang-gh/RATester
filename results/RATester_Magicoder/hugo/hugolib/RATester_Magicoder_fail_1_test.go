package hugolib

import (
	"fmt"
	"sync/atomic"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func Testfail_1(t *testing.T) {
	type fields struct {
		po                     *pageOutput
		otherOutputs           *maps.Cache[uint64, *pageContentOutput]
		contentRenderedVersion uint32
		contentRendered        atomic.Bool
		renderHooks            *renderHooks
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			pco := &pageContentOutput{
				po:                     tt.fields.po,
				otherOutputs:           tt.fields.otherOutputs,
				contentRenderedVersion: tt.fields.contentRenderedVersion,
				contentRendered:        tt.fields.contentRendered,
				renderHooks:            tt.fields.renderHooks,
			}
			pco.fail(tt.args.err)
		})
	}
}

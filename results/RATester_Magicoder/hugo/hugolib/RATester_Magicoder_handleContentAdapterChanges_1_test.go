package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugolib/pagesfromdata"
)

func TesthandleContentAdapterChanges_1(t *testing.T) {
	type args struct {
		bi          pagesfromdata.BuildInfo
		buildConfig *BuildCfg
	}
	tests := []struct {
		name string
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

			s := &Site{}
			s.handleContentAdapterChanges(tt.args.bi, tt.args.buildConfig)
		})
	}
}

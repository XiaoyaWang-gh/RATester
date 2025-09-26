package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
)

func TestCheckCascadePattern_1(t *testing.T) {
	type args struct {
		logger loggers.Logger
		m      PageMatcher
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

			CheckCascadePattern(tt.args.logger, tt.args.m)
		})
	}
}

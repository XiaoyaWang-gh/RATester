package langs

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestSetParams_1(t *testing.T) {
	type args struct {
		l      *Language
		params maps.Params
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

			SetParams(tt.args.l, tt.args.params)
		})
	}
}

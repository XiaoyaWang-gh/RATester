package page

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/types"
)

func TestexpandSummaryDivider_1(t *testing.T) {
	type args struct {
		s       string
		re      tagReStartEnd
		divider types.LowHigh[string]
	}
	tests := []struct {
		name  string
		args  args
		want  types.LowHigh[string]
		want1 types.LowHigh[string]
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

			got, got1 := expandSummaryDivider(tt.args.s, tt.args.re, tt.args.divider)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expandSummaryDivider() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("expandSummaryDivider() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

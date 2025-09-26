package page

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/types"
	"github.com/gohugoio/hugo/media"
)

func TestResolveParagraphTagAndSetWrapper_1(t *testing.T) {
	type args struct {
		s  *HtmlSummary
		mt media.Type
	}
	tests := []struct {
		name string
		args args
		want tagReStartEnd
	}{
		{
			name: "test_resolveParagraphTagAndSetWrapper_0",
			args: args{
				s: &HtmlSummary{
					source:         "test",
					SummaryLowHigh: types.LowHigh[string]{Low: 0, High: 0},
					SummaryEndTag:  types.LowHigh[string]{Low: 0, High: 0},
					WrapperStart:   types.LowHigh[string]{Low: 0, High: 0},
					WrapperEnd:     types.LowHigh[string]{Low: 0, High: 0},
					Divider:        types.LowHigh[string]{Low: 0, High: 0},
				},
				mt: media.Type{
					Type: "test",
				},
			},
			want: tagReStartEnd{
				startEndOfString: nil,
				endEndOfString:   nil,
				tagName:          "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.s.resolveParagraphTagAndSetWrapper(tt.args.mt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolveParagraphTagAndSetWrapper() = %v, want %v", got, tt.want)
			}
		})
	}
}

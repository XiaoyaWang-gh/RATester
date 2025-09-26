package page

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/media"
)

func TestExtractSummaryFromHTML_1(t *testing.T) {
	type args struct {
		mt       media.Type
		input    string
		numWords int
		isCJK    bool
	}
	tests := []struct {
		name string
		args args
		want HtmlSummary
	}{
		{
			name: "test case 1",
			args: args{
				mt:       media.Type{},
				input:    "",
				numWords: 0,
				isCJK:    false,
			},
			want: HtmlSummary{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ExtractSummaryFromHTML(tt.args.mt, tt.args.input, tt.args.numWords, tt.args.isCJK); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractSummaryFromHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}

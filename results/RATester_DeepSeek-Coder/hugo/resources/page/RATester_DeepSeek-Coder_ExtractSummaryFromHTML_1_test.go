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
		name       string
		args       args
		wantResult HtmlSummary
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

			gotResult := ExtractSummaryFromHTML(tt.args.mt, tt.args.input, tt.args.numWords, tt.args.isCJK)
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ExtractSummaryFromHTML() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

package page

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/media"
)

func TestExtractSummaryFromHTMLWithDivider_1(t *testing.T) {
	type args struct {
		mt      media.Type
		input   string
		divider string
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

			if gotResult := ExtractSummaryFromHTMLWithDivider(tt.args.mt, tt.args.input, tt.args.divider); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ExtractSummaryFromHTMLWithDivider() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

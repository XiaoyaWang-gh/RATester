package page

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/output"
)

func TestNewOutputFormat_1(t *testing.T) {
	type args struct {
		relPermalink string
		permalink    string
		isCanonical  bool
		f            output.Format
	}
	tests := []struct {
		name string
		args args
		want OutputFormat
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

			got := NewOutputFormat(tt.args.relPermalink, tt.args.permalink, tt.args.isCanonical, tt.args.f)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOutputFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

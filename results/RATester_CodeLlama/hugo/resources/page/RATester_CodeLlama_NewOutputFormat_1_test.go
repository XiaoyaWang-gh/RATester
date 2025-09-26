package page

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/output"
)

func TestNewOutputFormat_1(t *testing.T) {
	tests := []struct {
		name         string
		relPermalink string
		permalink    string
		isCanonical  bool
		f            output.Format
		want         OutputFormat
	}{
		{
			name:         "html",
			relPermalink: "relPermalink",
			permalink:    "permalink",
			isCanonical:  true,
			f: output.Format{
				Name: "html",
				Rel:  "rel",
			},
			want: OutputFormat{
				Rel: "canonical",
				Format: output.Format{
					Name: "html",
					Rel:  "rel",
				},
				relPermalink: "relPermalink",
				permalink:    "permalink",
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

			if got := NewOutputFormat(tt.relPermalink, tt.permalink, tt.isCanonical, tt.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOutputFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

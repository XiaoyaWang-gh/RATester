package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/media"
	"github.com/gohugoio/hugo/output"
)

func TestMediaType_5(t *testing.T) {
	type fields struct {
		Rel          string
		Format       output.Format
		relPermalink string
		permalink    string
	}
	tests := []struct {
		name   string
		fields fields
		want   media.Type
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

			o := OutputFormat{
				Rel:          tt.fields.Rel,
				Format:       tt.fields.Format,
				relPermalink: tt.fields.relPermalink,
				permalink:    tt.fields.permalink,
			}
			if got := o.MediaType(); got != tt.want {
				t.Errorf("OutputFormat.MediaType() = %v, want %v", got, tt.want)
			}
		})
	}
}

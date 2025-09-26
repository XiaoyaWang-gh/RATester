package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/output"
)

func TestName_10(t *testing.T) {
	type fields struct {
		Rel          string
		Format       output.Format
		relPermalink string
		permalink    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				Rel: "canonical",
				Format: output.Format{
					Name: "HTML",
				},
				relPermalink: "",
				permalink:    "",
			},
			want: "HTML",
		},
		{
			name: "Test Case 2",
			fields: fields{
				Rel: "amphtml",
				Format: output.Format{
					Name: "AMP",
				},
				relPermalink: "",
				permalink:    "",
			},
			want: "AMP",
		},
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
			if got := o.Name(); got != tt.want {
				t.Errorf("OutputFormat.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

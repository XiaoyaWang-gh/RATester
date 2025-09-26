package page

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/media"
	"github.com/gohugoio/hugo/output"
)

func TestMediaType_5(t *testing.T) {
	tests := []struct {
		name string
		o    OutputFormat
		want media.Type
	}{
		{
			name: "Test case 1",
			o: OutputFormat{
				Format: output.Format{
					MediaType: media.Type{
						Type: "application/rss+xml",
					},
				},
			},
			want: media.Type{
				Type: "application/rss+xml",
			},
		},
		{
			name: "Test case 2",
			o: OutputFormat{
				Format: output.Format{
					MediaType: media.Type{
						Type: "text/html",
					},
				},
			},
			want: media.Type{
				Type: "text/html",
			},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.o.MediaType(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MediaType() = %v, want %v", got, tt.want)
			}
		})
	}
}

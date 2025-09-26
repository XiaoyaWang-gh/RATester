package urls

import (
	"fmt"
	"net/url"
	"testing"
)

func TestString_23(t *testing.T) {
	type fields struct {
		url                     *url.URL
		WithPath                string
		WithPathNoTrailingSlash string
		WithoutPath             string
		BasePath                string
		BasePathNoTrailingSlash string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
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

			b := BaseURL{
				url:                     tt.fields.url,
				WithPath:                tt.fields.WithPath,
				WithPathNoTrailingSlash: tt.fields.WithPathNoTrailingSlash,
				WithoutPath:             tt.fields.WithoutPath,
				BasePath:                tt.fields.BasePath,
				BasePathNoTrailingSlash: tt.fields.BasePathNoTrailingSlash,
			}
			if got := b.String(); got != tt.want {
				t.Errorf("BaseURL.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

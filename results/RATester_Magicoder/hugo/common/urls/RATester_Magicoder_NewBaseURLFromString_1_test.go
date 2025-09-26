package urls

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestNewBaseURLFromString_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    BaseURL
		wantErr bool
	}{
		{
			name:  "valid url",
			input: "http://example.com",
			want: BaseURL{
				url:                     &url.URL{Scheme: "http", Host: "example.com"},
				WithPath:                "http://example.com",
				WithPathNoTrailingSlash: "http://example.com",
				WithoutPath:             "http://example.com",
				BasePath:                "/",
				BasePathNoTrailingSlash: "/",
			},
			wantErr: false,
		},
		{
			name:    "invalid url",
			input:   "://example.com",
			want:    BaseURL{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := NewBaseURLFromString(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBaseURLFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBaseURLFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

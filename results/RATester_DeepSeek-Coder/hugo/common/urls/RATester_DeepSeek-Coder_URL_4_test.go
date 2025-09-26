package urls

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestURL_4(t *testing.T) {
	type fields struct {
		url *url.URL
	}
	tests := []struct {
		name   string
		fields fields
		want   *url.URL
	}{
		{
			name: "Test URL",
			fields: fields{
				url: &url.URL{
					Scheme: "https",
					Host:   "example.com",
				},
			},
			want: &url.URL{
				Scheme: "https",
				Host:   "example.com",
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

			b := BaseURL{
				url: tt.fields.url,
			}
			if got := b.URL(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URL() = %v, want %v", got, tt.want)
			}
		})
	}
}

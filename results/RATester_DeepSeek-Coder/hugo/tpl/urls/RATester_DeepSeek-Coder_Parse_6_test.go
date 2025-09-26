package urls

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestParse_6(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		rawurl  any
		want    *url.URL
		wantErr bool
	}{
		{
			name:   "valid url",
			rawurl: "https://example.com",
			want: &url.URL{
				Scheme: "https",
				Host:   "example.com",
			},
			wantErr: false,
		},
		{
			name:    "invalid url",
			rawurl:  ":",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "nil url",
			rawurl:  nil,
			want:    nil,
			wantErr: true,
		},
	}

	ns := &Namespace{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.Parse(tt.rawurl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

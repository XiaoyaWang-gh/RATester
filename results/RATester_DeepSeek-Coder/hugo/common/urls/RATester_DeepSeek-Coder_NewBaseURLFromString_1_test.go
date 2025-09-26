package urls

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestNewBaseURLFromString_1(t *testing.T) {
	type args struct {
		b string
	}
	tests := []struct {
		name    string
		args    args
		want    BaseURL
		wantErr bool
	}{
		{
			name: "valid url",
			args: args{
				b: "https://example.com",
			},
			want: BaseURL{
				url: &url.URL{
					Scheme: "https",
					Host:   "example.com",
				},
				WithPath:                "https://example.com/",
				WithPathNoTrailingSlash: "https://example.com",
				WithoutPath:             "https://example.com",
				BasePath:                "/",
				BasePathNoTrailingSlash: "",
			},
			wantErr: false,
		},
		{
			name: "invalid url",
			args: args{
				b: ":",
			},
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

			got, err := NewBaseURLFromString(tt.args.b)
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

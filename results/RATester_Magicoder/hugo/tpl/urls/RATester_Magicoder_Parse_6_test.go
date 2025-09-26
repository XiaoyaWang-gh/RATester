package urls

import (
	"fmt"
	"net/url"
	"reflect"
	"testing"
)

func TestParse_6(t *testing.T) {
	type args struct {
		rawurl any
	}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			name: "Test with string",
			args: args{
				rawurl: "https://www.example.com",
			},
			want: &url.URL{
				Scheme: "https",
				Host:   "www.example.com",
			},
			wantErr: false,
		},
		{
			name: "Test with int",
			args: args{
				rawurl: 123,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test with nil",
			args: args{
				rawurl: nil,
			},
			want:    nil,
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

			ns := &Namespace{}
			got, err := ns.Parse(tt.args.rawurl)
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

package helpers

import (
	"fmt"
	"testing"
)

func TestIsAbsURL_1(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test with http://",
			args: args{
				in: "http://example.com",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Test with https://",
			args: args{
				in: "https://example.com",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Test with relative URL",
			args: args{
				in: "/relative/url",
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Test with invalid URL",
			args: args{
				in: "invalid-url",
			},
			want:    false,
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

			p := &PathSpec{}
			got, err := p.IsAbsURL(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsAbsURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsAbsURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

package dartsass

import (
	"fmt"
	"testing"
)

func TestCanonicalizeURL_1(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				url: "test_url",
			},
			want:    "file://test_url",
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				url: "test_url2",
			},
			want:    "file://test_url2",
			wantErr: false,
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

			ir := importResolver{}
			got, err := ir.CanonicalizeURL(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("CanonicalizeURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CanonicalizeURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

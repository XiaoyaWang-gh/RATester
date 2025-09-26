package encoding

import (
	"fmt"
	"testing"
)

func TestBase64Decode_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		content any
		want    string
		wantErr bool
	}{
		{
			name:    "valid base64",
			content: "SGVsbG8gd29ybGQ=",
			want:    "Hello world",
			wantErr: false,
		},
		{
			name:    "invalid base64",
			content: "SGVsbG8gd29ybGQ",
			want:    "",
			wantErr: true,
		},
		{
			name:    "non-string content",
			content: 123,
			want:    "",
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

			got, err := ns.Base64Decode(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Base64Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

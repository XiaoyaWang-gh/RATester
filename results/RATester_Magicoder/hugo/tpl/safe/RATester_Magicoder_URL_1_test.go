package safe

import (
	"fmt"
	"html/template"
	"testing"
)

func TestURL_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    template.URL
		wantErr bool
	}{
		{
			name:  "valid string",
			input: "https://example.com",
			want:  "https://example.com",
		},
		{
			name:    "invalid type",
			input:   123,
			wantErr: true,
		},
		{
			name:    "nil",
			input:   nil,
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

			got, err := ns.URL(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("URL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("URL() = %v, want %v", got, tt.want)
			}
		})
	}
}

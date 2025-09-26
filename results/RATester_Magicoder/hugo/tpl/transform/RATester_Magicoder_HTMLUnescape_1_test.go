package transform

import (
	"fmt"
	"testing"
)

func TestHTMLUnescape_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:  "valid input",
			input: "&lt;p&gt;Hello, World!&lt;/p&gt;",
			want:  "<p>Hello, World!</p>",
		},
		{
			name:    "invalid input",
			input:   "<p>Hello, World!</p>",
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

			got, err := ns.HTMLUnescape(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTMLUnescape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("HTMLUnescape() = %v, want %v", got, tt.want)
			}
		})
	}
}

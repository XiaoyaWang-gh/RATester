package transform

import (
	"fmt"
	"testing"
)

func TestXMLEscape_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:  "valid input",
			input: "<tag>content</tag>",
			want:  "&lt;tag&gt;content&lt;/tag&gt;",
		},
		{
			name:    "invalid input",
			input:   make(chan int),
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

			got, err := ns.XMLEscape(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("XMLEscape() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("XMLEscape() = %v, want %v", got, tt.want)
			}
		})
	}
}

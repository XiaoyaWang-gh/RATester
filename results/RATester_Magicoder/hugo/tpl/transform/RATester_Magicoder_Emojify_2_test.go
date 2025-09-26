package transform

import (
	"fmt"
	"html/template"
	"testing"
)

func TestEmojify_2(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    template.HTML
		wantErr bool
	}{
		{
			name:  "valid input",
			input: ":smile:",
			want:  "😄",
		},
		{
			name:    "invalid input",
			input:   make(chan int),
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

			got, err := ns.Emojify(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Emojify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Emojify() = %v, want %v", got, tt.want)
			}
		})
	}
}

package transform

import (
	"fmt"
	"testing"
)

func TeststringToRune_1(t *testing.T) {
	tests := []struct {
		name    string
		input   any
		want    rune
		wantErr bool
	}{
		{
			name:    "valid string",
			input:   "a",
			want:    'a',
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   "",
			want:    0,
			wantErr: false,
		},
		{
			name:    "invalid input",
			input:   123,
			want:    0,
			wantErr: true,
		},
		{
			name:    "multiple characters",
			input:   "ab",
			want:    0,
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

			got, err := stringToRune(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("stringToRune() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("stringToRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

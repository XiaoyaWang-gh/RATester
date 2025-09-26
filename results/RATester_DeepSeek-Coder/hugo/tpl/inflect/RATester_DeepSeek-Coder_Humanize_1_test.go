package inflect

import (
	"fmt"
	"testing"
)

func TestHumanize_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	tests := []struct {
		name    string
		input   any
		want    string
		wantErr bool
	}{
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "non-empty string",
			input: "test",
			want:  "Test",
		},
		{
			name:  "integer",
			input: 1,
			want:  "1st",
		},
		{
			name:    "float",
			input:   1.1,
			want:    "",
			wantErr: true,
		},
		{
			name:    "boolean",
			input:   true,
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

			got, err := ns.Humanize(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Namespace.Humanize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Namespace.Humanize() = %v, want %v", got, tt.want)
			}
		})
	}
}

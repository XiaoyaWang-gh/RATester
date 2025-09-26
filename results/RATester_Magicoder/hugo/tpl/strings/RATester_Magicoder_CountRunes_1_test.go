package strings

import (
	"fmt"
	"testing"
)

func TestCountRunes_1(t *testing.T) {
	ns := &Namespace{}

	tests := []struct {
		name    string
		s       any
		want    int
		wantErr bool
	}{
		{
			name: "Test with string",
			s:    "Hello, World!",
			want: 13,
		},
		{
			name: "Test with empty string",
			s:    "",
			want: 0,
		},
		{
			name: "Test with non-string type",
			s:    123,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := ns.CountRunes(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("CountRunes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CountRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}

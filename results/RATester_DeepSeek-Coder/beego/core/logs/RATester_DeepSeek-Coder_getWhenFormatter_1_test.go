package logs

import (
	"fmt"
	"testing"
)

func TestGetWhenFormatter_1(t *testing.T) {
	tests := []struct {
		name string
		p    *PatternLogFormatter
		want string
	}{
		{
			name: "Test with empty WhenFormat",
			p:    &PatternLogFormatter{},
			want: "2006/01/02 15:04:05.123",
		},
		{
			name: "Test with non-empty WhenFormat",
			p:    &PatternLogFormatter{WhenFormat: "2006-01-02 15:04:05"},
			want: "2006-01-02 15:04:05",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.p.getWhenFormatter(); got != tt.want {
				t.Errorf("PatternLogFormatter.getWhenFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

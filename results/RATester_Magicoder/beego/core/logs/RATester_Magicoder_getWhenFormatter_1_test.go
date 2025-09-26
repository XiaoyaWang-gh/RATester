package logs

import (
	"fmt"
	"testing"
)

func TestgetWhenFormatter_1(t *testing.T) {
	tests := []struct {
		name string
		p    *PatternLogFormatter
		want string
	}{
		{
			name: "Test Case 1",
			p:    &PatternLogFormatter{WhenFormat: "2006-01-02"},
			want: "2006-01-02",
		},
		{
			name: "Test Case 2",
			p:    &PatternLogFormatter{},
			want: "2006/01/02 15:04:05.123",
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
				t.Errorf("getWhenFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

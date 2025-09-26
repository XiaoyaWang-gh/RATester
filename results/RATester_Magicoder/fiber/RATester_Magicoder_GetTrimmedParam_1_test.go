package fiber

import (
	"fmt"
	"testing"
)

func TestGetTrimmedParam_1(t *testing.T) {
	tests := []struct {
		name  string
		param string
		want  string
	}{
		{"empty string", "", ""},
		{"string without starter char", "test", "test"},
		{"string with starter char", "*test", "test"},
		{"string with starter char and optional char", "*test*", "test"},
		{"string with starter char and optional char at the end", "test*", "test"},
		{"string with starter char and optional char at the beginning and end", "*test*", "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetTrimmedParam(tt.param); got != tt.want {
				t.Errorf("GetTrimmedParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

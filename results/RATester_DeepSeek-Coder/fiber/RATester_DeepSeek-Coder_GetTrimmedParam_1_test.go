package fiber

import (
	"fmt"
	"testing"
)

func TestGetTrimmedParam_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"Empty string", "", ""},
		{"String without starter char", "test", "test"},
		{"String with starter char", "*test", "test"},
		{"String with starter char at the end", "test*", "test"},
		{"String with starter char at both ends", "*test*", "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetTrimmedParam(tt.arg); got != tt.want {
				t.Errorf("GetTrimmedParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

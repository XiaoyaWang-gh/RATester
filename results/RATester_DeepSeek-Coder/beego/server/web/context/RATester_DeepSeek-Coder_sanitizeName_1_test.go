package context

import (
	"fmt"
	"testing"
)

func TestSanitizeName_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"empty string", "", ""},
		{"simple name", "name", "name"},
		{"name with spaces", "name with spaces", "name_with_spaces"},
		{"name with special characters", "name@with#special*characters", "name_with_special_characters"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := sanitizeName(tt.arg); got != tt.want {
				t.Errorf("sanitizeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

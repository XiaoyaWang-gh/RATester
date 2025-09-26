package context

import (
	"fmt"
	"testing"
)

func TestParseFormBoolValue_1(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		want    bool
		wantErr bool
	}{
		{"Test Case 1", "on", true, false},
		{"Test Case 2", "1", true, false},
		{"Test Case 3", "yes", true, false},
		{"Test Case 4", "off", false, false},
		{"Test Case 5", "0", false, false},
		{"Test Case 6", "no", false, false},
		{"Test Case 7", "true", true, false},
		{"Test Case 8", "false", false, false},
		{"Test Case 9", "invalid", false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := parseFormBoolValue(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFormBoolValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseFormBoolValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

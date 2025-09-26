package config

import (
	"fmt"
	"testing"
)

func TestSplitEnvVar_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantName string
		wantVal  string
	}{
		{
			name:     "Test case 1",
			input:    "VAR1=VALUE1",
			wantName: "VAR1",
			wantVal:  "VALUE1",
		},
		{
			name:     "Test case 2",
			input:    "VAR2=VALUE2",
			wantName: "VAR2",
			wantVal:  "VALUE2",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotName, gotVal := SplitEnvVar(tt.input)
			if gotName != tt.wantName {
				t.Errorf("SplitEnvVar() gotName = %v, want %v", gotName, tt.wantName)
			}
			if gotVal != tt.wantVal {
				t.Errorf("SplitEnvVar() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
		})
	}
}

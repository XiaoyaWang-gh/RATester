package config

import (
	"fmt"
	"testing"
)

func TestDetectLegacyINIFormat_1(t *testing.T) {
	tests := []struct {
		name     string
		content  []byte
		expected bool
	}{
		{
			name:     "common section",
			content:  []byte("[common]\n"),
			expected: true,
		},
		{
			name:     "no common section",
			content:  []byte("[other]\n"),
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := DetectLegacyINIFormat(tt.content); got != tt.expected {
				t.Errorf("DetectLegacyINIFormat() = %v, want %v", got, tt.expected)
			}
		})
	}
}

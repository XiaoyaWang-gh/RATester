package exif

import (
	"fmt"
	"regexp"
	"testing"
)

func TestshouldInclude_1(t *testing.T) {
	d := &Decoder{
		includeFieldsRe: regexp.MustCompile("^(GPS|EXIF|IPTC)$"),
	}

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Include", "GPS", true},
		{"Exclude", "JPEG", false},
		{"Include", "EXIF", true},
		{"Exclude", "IPTC", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := d.shouldInclude(test.input)
			if result != test.expected {
				t.Errorf("Expected %t, but got %t", test.expected, result)
			}
		})
	}
}

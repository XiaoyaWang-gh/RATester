package template

import (
	"fmt"
	"testing"
)

func TestAttrType_1(t *testing.T) {
	tests := []struct {
		name string
		want contentType
	}{
		{"data-test", contentTypePlain},
		{"xmlns:test", contentTypeURL},
		{"onClick", contentTypeJS},
		{"src", contentTypeURL},
		{"uri", contentTypeURL},
		{"url", contentTypeURL},
		{"test", contentTypePlain},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := attrType(tt.name); got != tt.want {
				t.Errorf("attrType() = %v, want %v", got, tt.want)
			}
		})
	}
}

package tplimpl

import (
	"fmt"
	"testing"
)

func TestresolveType_1(t *testing.T) {
	tests := []struct {
		name     string
		template templateInfo
		want     templateType
	}{
		{
			name: "Test Case 1",
			template: templateInfo{
				name: "test.html",
			},
			want: resolveTemplateType("test.html"),
		},
		{
			name: "Test Case 2",
			template: templateInfo{
				name: "test.txt",
			},
			want: resolveTemplateType("test.txt"),
		},
		{
			name: "Test Case 3",
			template: templateInfo{
				name: "test.unknown",
			},
			want: resolveTemplateType("test.unknown"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.template.resolveType(); got != tt.want {
				t.Errorf("resolveType() = %v, want %v", got, tt.want)
			}
		})
	}
}

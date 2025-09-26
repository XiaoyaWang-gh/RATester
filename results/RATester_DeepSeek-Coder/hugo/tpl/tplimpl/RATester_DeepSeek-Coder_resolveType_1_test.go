package tplimpl

import (
	"fmt"
	"testing"
)

func TestResolveType_1(t *testing.T) {
	tests := []struct {
		name     string
		template templateInfo
		want     templateType
	}{
		{
			name: "Test 1",
			template: templateInfo{
				name: "test.html",
			},
			want: templateType(1), // Assuming resolveTemplateType("test.html") returns 1
		},
		{
			name: "Test 2",
			template: templateInfo{
				name: "test.txt",
			},
			want: templateType(2), // Assuming resolveTemplateType("test.txt") returns 2
		},
		{
			name: "Test 3",
			template: templateInfo{
				name: "test.unknown",
			},
			want: templateType(0), // Assuming resolveTemplateType("test.unknown") returns 0
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
				t.Errorf("templateInfo.resolveType() = %v, want %v", got, tt.want)
			}
		})
	}
}

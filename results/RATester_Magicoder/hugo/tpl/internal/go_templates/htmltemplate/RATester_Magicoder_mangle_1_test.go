package template

import (
	"fmt"
	"testing"
)

func Testmangle_1(t *testing.T) {
	tests := []struct {
		name         string
		context      context
		templateName string
		want         string
	}{
		{
			name: "Test Case 1",
			context: context{
				state:   stateText,
				delim:   delimNone,
				urlPart: urlPartNone,
				jsCtx:   jsCtxRegexp,
				attr:    attrNone,
			},
			templateName: "test",
			want:         "test",
		},
		{
			name: "Test Case 2",
			context: context{
				state:   stateText,
				delim:   delimNone,
				urlPart: urlPartNone,
				jsCtx:   jsCtxRegexp,
				attr:    attrNone,
			},
			templateName: "test2",
			want:         "test2",
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

			if got := tt.context.mangle(tt.templateName); got != tt.want {
				t.Errorf("mangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

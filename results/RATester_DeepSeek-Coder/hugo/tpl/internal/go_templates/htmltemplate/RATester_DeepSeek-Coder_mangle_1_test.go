package template

import (
	"fmt"
	"testing"
)

func TestMangle_1(t *testing.T) {
	tests := []struct {
		name         string
		c            context
		templateName string
		want         string
	}{
		{
			name: "Test case 1",
			c: context{
				state:   stateText,
				delim:   delimNone,
				urlPart: urlPartNone,
				jsCtx:   jsCtxRegexp,
				attr:    attrNone,
				element: elementNone,
			},
			templateName: "test",
			want:         "test",
		},
		{
			name: "Test case 2",
			c: context{
				state:   stateText,
				delim:   delimNone,
				urlPart: urlPartNone,
				jsCtx:   jsCtxRegexp,
				attr:    attrNone,
				element: elementNone,
			},
			templateName: "test2",
			want:         "test2$htmltemplate_0_0_0_0_0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.mangle(tt.templateName); got != tt.want {
				t.Errorf("context.mangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

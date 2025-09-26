package tplimpl

import (
	"fmt"
	"testing"
)

func TestIsInternalTemplate_1(t *testing.T) {
	testCases := []struct {
		name     string
		template *templateState
		want     bool
	}{
		{
			name: "Internal template",
			template: &templateState{
				info: templateInfo{
					isEmbedded: true,
				},
			},
			want: true,
		},
		{
			name: "External template",
			template: &templateState{
				info: templateInfo{
					isEmbedded: false,
				},
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tc.template.IsInternalTemplate()
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

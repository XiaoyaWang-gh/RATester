package template

import (
	"fmt"
	"testing"
)

func TestCheckCanParse_1(t *testing.T) {
	type testCase struct {
		name     string
		template *Template
		wantErr  bool
	}

	testCases := []testCase{
		{
			name:     "NilTemplate",
			template: nil,
			wantErr:  false,
		},
		{
			name:     "EscapedTemplate",
			template: &Template{nameSpace: &nameSpace{escaped: true}},
			wantErr:  true,
		},
		{
			name:     "NotEscapedTemplate",
			template: &Template{nameSpace: &nameSpace{escaped: false}},
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tc.template.checkCanParse()
			if (err != nil) != tc.wantErr {
				t.Errorf("checkCanParse() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

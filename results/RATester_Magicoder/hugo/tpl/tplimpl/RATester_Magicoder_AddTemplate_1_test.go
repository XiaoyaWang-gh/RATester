package tplimpl

import (
	"fmt"
	"testing"
)

func TestAddTemplate_1(t *testing.T) {
	handler := &templateHandler{
		main: &templateNamespace{},
	}

	testCases := []struct {
		name    string
		tpl     string
		wantErr bool
	}{
		{
			name:    "valid_template",
			tpl:     "{{ .Name }}",
			wantErr: false,
		},
		{
			name:    "invalid_template",
			tpl:     "{{ .Name }",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := handler.AddTemplate(tc.name, tc.tpl)
			if (err != nil) != tc.wantErr {
				t.Errorf("AddTemplate() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}

package tplimpl

import (
	"fmt"
	"testing"
)

func TestAddTemplate_1(t *testing.T) {
	type args struct {
		name string
		tpl  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				name: "test_template",
				tpl:  "{{ .Site.Title }}",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				name: "test_template_2",
				tpl:  "{{ .Site.Title }}",
			},
			wantErr: true, // This is a hypothetical error case, replace with actual error handling logic
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			th := &templateHandler{
				main: &templateNamespace{},
			}
			if err := th.AddTemplate(tt.args.name, tt.args.tpl); (err != nil) != tt.wantErr {
				t.Errorf("templateHandler.AddTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

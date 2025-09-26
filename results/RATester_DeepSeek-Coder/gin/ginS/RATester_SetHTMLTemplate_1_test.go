package ginS

import (
	"fmt"
	"html/template"
	"testing"
)

func TestSetHTMLTemplate_1(t *testing.T) {
	type args struct {
		templ *template.Template
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				templ: template.Must(template.New("test").Parse("test")),
			},
		},
		// Add more test cases here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			SetHTMLTemplate(tt.args.templ)
		})
	}
}

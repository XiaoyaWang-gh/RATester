package gin

import (
	"fmt"
	"html/template"
	"testing"
)

func TestDebugPrintLoadTemplate_1(t *testing.T) {
	type args struct {
		tmpl *template.Template
	}

	// Create a new template with a name and content
	testTemplate := template.New("testTemplate")
	testTemplate, err := testTemplate.Parse("This is a test template")
	if err != nil {
		t.Errorf("Failed to parse test template: %v", err)
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test with a valid template",
			args: args{tmpl: testTemplate},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			debugPrintLoadTemplate(tt.args.tmpl)
		})
	}
}

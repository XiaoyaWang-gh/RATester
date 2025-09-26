package render

import (
	"fmt"
	"html/template"
	"reflect"
	"testing"
)

func TestLoadTemplate_1(t *testing.T) {
	type fields struct {
		Files   []string
		Glob    string
		Delims  Delims
		FuncMap template.FuncMap
	}
	tests := []struct {
		name   string
		fields fields
		want   *template.Template
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			h := &HTMLDebug{
				Files:   tt.fields.Files,
				Glob:    tt.fields.Glob,
				Delims:  tt.fields.Delims,
				FuncMap: tt.fields.FuncMap,
			}
			if got := h.loadTemplate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

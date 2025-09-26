package render

import (
	"fmt"
	"html/template"
	"reflect"
	"testing"
)

func TestloadTemplate_1(t *testing.T) {
	tests := []struct {
		name string
		r    HTMLDebug
		want *template.Template
	}{
		{
			name: "Test with files",
			r: HTMLDebug{
				Files: []string{"file1.html", "file2.html"},
			},
			want: template.Must(template.New("").ParseFiles("file1.html", "file2.html")),
		},
		{
			name: "Test with glob",
			r: HTMLDebug{
				Glob: "*.html",
			},
			want: template.Must(template.New("").ParseGlob("*.html")),
		},
		{
			name: "Test without files or glob",
			r:    HTMLDebug{},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.r.loadTemplate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}

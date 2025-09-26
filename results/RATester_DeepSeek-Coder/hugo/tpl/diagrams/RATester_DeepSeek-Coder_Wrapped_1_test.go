package diagrams

import (
	"fmt"
	"html/template"
	"testing"

	"github.com/bep/goat"
)

func TestWrapped_1(t *testing.T) {
	type fields struct {
		d goat.SVG
	}
	tests := []struct {
		name   string
		fields fields
		want   template.HTML
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

			d := goatDiagram{
				d: tt.fields.d,
			}
			if got := d.Wrapped(); got != tt.want {
				t.Errorf("goatDiagram.Wrapped() = %v, want %v", got, tt.want)
			}
		})
	}
}

package diagrams

import (
	"fmt"
	"html/template"
	"testing"

	"github.com/bep/goat"
)

func TestInner_2(t *testing.T) {
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
		d := goatDiagram{
			d: tt.fields.d,
		}
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := d.Inner(); got != tt.want {
				t.Errorf("goatDiagram.Inner() = %v, want %v", got, tt.want)
			}
		})
	}
}

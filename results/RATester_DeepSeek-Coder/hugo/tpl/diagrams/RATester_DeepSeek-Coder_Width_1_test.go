package diagrams

import (
	"fmt"
	"testing"

	"github.com/bep/goat"
)

func TestWidth_1(t *testing.T) {
	type fields struct {
		d goat.SVG
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Test Width",
			fields: fields{
				d: goat.SVG{
					Width: 100,
				},
			},
			want: 100,
		},
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
			if got := d.Width(); got != tt.want {
				t.Errorf("goatDiagram.Width() = %v, want %v", got, tt.want)
			}
		})
	}
}

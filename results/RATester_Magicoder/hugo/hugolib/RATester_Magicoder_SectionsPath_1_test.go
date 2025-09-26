package hugolib

import (
	"fmt"
	"testing"
)

func TestSectionsPath_1(t *testing.T) {
	type fields struct {
		p *pageState
	}
	tests := []struct {
		name   string
		fields fields
		want   string
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

			p := pageTree{
				p: tt.fields.p,
			}
			if got := p.SectionsPath(); got != tt.want {
				t.Errorf("SectionsPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

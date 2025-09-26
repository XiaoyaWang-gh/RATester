package hugolib

import (
	"fmt"
	"testing"
)

func TestLang_3(t *testing.T) {
	type fields struct {
		s *Site
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

			p := &pageMeta{
				s: tt.fields.s,
			}
			if got := p.Lang(); got != tt.want {
				t.Errorf("pageMeta.Lang() = %v, want %v", got, tt.want)
			}
		})
	}
}

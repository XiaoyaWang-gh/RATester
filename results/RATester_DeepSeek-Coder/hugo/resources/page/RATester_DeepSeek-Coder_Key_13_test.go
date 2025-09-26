package page

import (
	"fmt"
	"testing"
)

func TestKey_13(t *testing.T) {
	type fields struct {
		s Site
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

			s := &siteWrapper{
				s: tt.fields.s,
			}
			if got := s.Key(); got != tt.want {
				t.Errorf("siteWrapper.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}

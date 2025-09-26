package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/hugo"
	"github.com/gohugoio/hugo/langs"
)

func TestLanguageCode_3(t *testing.T) {
	type fields struct {
		h hugo.HugoInfo
		l *langs.Language
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

			s := testSite{
				h: tt.fields.h,
				l: tt.fields.l,
			}
			if got := s.LanguageCode(); got != tt.want {
				t.Errorf("LanguageCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

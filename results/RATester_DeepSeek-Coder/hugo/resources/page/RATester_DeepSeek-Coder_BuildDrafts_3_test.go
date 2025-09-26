package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/hugo"
	"github.com/gohugoio/hugo/langs"
)

func TestBuildDrafts_3(t *testing.T) {
	type fields struct {
		h hugo.HugoInfo
		l *langs.Language
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test BuildDrafts",
			fields: fields{
				h: hugo.HugoInfo{},
				l: &langs.Language{},
			},
			want: false,
		},
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
			if got := s.BuildDrafts(); got != tt.want {
				t.Errorf("testSite.BuildDrafts() = %v, want %v", got, tt.want)
			}
		})
	}
}

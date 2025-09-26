package langs

import (
	"fmt"
	"testing"
	"time"

	"github.com/gohugoio/hugo/common/htime"
	"github.com/gohugoio/hugo/common/maps"
	"github.com/gohugoio/locales"
	"golang.org/x/text/language"
)

func TestString_5(t *testing.T) {
	type fields struct {
		Lang           string
		LanguageConfig LanguageConfig
		translator     locales.Translator
		timeFormatter  htime.TimeFormatter
		tag            language.Tag
		collator1      *Collator
		collator2      *Collator
		location       *time.Location
		params         maps.Params
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				Lang: "en",
			},
			want: "en",
		},
		{
			name: "Test Case 2",
			fields: fields{
				Lang: "fr",
			},
			want: "fr",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &Language{
				Lang:           tt.fields.Lang,
				LanguageConfig: tt.fields.LanguageConfig,
				translator:     tt.fields.translator,
				timeFormatter:  tt.fields.timeFormatter,
				tag:            tt.fields.tag,
				collator1:      tt.fields.collator1,
				collator2:      tt.fields.collator2,
				location:       tt.fields.location,
				params:         tt.fields.params,
			}
			if got := l.String(); got != tt.want {
				t.Errorf("Language.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

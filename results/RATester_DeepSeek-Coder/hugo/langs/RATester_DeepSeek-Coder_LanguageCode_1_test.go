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

func TestLanguageCode_1(t *testing.T) {
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
			name: "Test LanguageCode with LanguageConfig.LanguageCode not empty",
			fields: fields{
				Lang: "en",
				LanguageConfig: LanguageConfig{
					LanguageCode: "en-US",
				},
			},
			want: "en-US",
		},
		{
			name: "Test LanguageCode with LanguageConfig.LanguageCode empty",
			fields: fields{
				Lang: "en",
				LanguageConfig: LanguageConfig{
					LanguageCode: "",
				},
			},
			want: "en",
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
			if got := l.LanguageCode(); got != tt.want {
				t.Errorf("Language.LanguageCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

package langs

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/gohugoio/hugo/common/htime"
	"github.com/gohugoio/hugo/common/maps"
	"github.com/gohugoio/locales"
	"golang.org/x/text/language"
)

func TestParams_2(t *testing.T) {
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
		want   maps.Params
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
			if got := l.Params(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Language.Params() = %v, want %v", got, tt.want)
			}
		})
	}
}

package exif

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/bep/logg"
)

func TestShouldInclude_1(t *testing.T) {
	type fields struct {
		includeFieldsRe  *regexp.Regexp
		excludeFieldsrRe *regexp.Regexp
		noDate           bool
		noLatLong        bool
		warnl            logg.LevelLogger
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
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

			d := &Decoder{
				includeFieldsRe:  tt.fields.includeFieldsRe,
				excludeFieldsrRe: tt.fields.excludeFieldsrRe,
				noDate:           tt.fields.noDate,
				noLatLong:        tt.fields.noLatLong,
				warnl:            tt.fields.warnl,
			}
			if got := d.shouldInclude(tt.args.s); got != tt.want {
				t.Errorf("Decoder.shouldInclude() = %v, want %v", got, tt.want)
			}
		})
	}
}

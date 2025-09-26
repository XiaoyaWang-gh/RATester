package htime

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/locales"
)

func TestNewTimeFormatter_1(t *testing.T) {
	type args struct {
		ltr locales.Translator
	}
	tests := []struct {
		name string
		args args
		want TimeFormatter
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

			if got := NewTimeFormatter(tt.args.ltr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTimeFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

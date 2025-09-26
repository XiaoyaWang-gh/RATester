package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func Testc_1(t *testing.T) {
	type fields struct {
		po *pageOutput
	}
	tests := []struct {
		name   string
		fields fields
		want   page.Markup
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

			pco := &pageContentOutput{
				po: tt.fields.po,
			}
			if got := pco.c(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pageContentOutput.c() = %v, want %v", got, tt.want)
			}
		})
	}
}

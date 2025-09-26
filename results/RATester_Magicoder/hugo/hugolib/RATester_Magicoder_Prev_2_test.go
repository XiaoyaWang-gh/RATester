package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestPrev_2(t *testing.T) {
	type fields struct {
		nextPrev *nextPrev
	}
	tests := []struct {
		name   string
		fields fields
		want   page.Page
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

			p := pagePosition{
				nextPrev: tt.fields.nextPrev,
			}
			if got := p.Prev(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}

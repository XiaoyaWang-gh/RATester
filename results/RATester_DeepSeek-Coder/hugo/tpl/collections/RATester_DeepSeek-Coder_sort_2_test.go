package collections

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/langs"
	"github.com/gohugoio/hugo/tpl/compare"
)

func TestSort_2(t *testing.T) {
	type fields struct {
		Collator  *langs.Collator
		sortComp  *compare.Namespace
		Pairs     []pair
		SortAsc   bool
		SliceType reflect.Type
	}
	tests := []struct {
		name   string
		fields fields
		want   any
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

			p := pairList{
				Collator:  tt.fields.Collator,
				sortComp:  tt.fields.sortComp,
				Pairs:     tt.fields.Pairs,
				SortAsc:   tt.fields.SortAsc,
				SliceType: tt.fields.SliceType,
			}
			if got := p.sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pairList.sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

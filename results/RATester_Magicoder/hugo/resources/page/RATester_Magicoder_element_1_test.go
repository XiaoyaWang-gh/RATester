package page

import (
	"fmt"
	"reflect"
	"testing"
)

func Testelement_1(t *testing.T) {
	type fields struct {
		number    int
		Paginator *Paginator
	}
	tests := []struct {
		name   string
		fields fields
		want   paginatedElement
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

			p := &Pager{
				number:    tt.fields.number,
				Paginator: tt.fields.Paginator,
			}
			if got := p.element(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pager.element() = %v, want %v", got, tt.want)
			}
		})
	}
}

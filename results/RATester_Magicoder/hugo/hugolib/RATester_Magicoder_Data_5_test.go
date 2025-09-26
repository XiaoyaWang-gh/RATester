package hugolib

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestData_5(t *testing.T) {
	type fields struct {
		pageState *pageState
		dataInit  sync.Once
		data      page.Data
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

			p := &pageData{
				pageState: tt.fields.pageState,
				dataInit:  tt.fields.dataInit,
				data:      tt.fields.data,
			}
			if got := p.Data(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pageData.Data() = %v, want %v", got, tt.want)
			}
		})
	}
}

package pagination

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIsActive_1(t *testing.T) {
	type fields struct {
		Request     *http.Request
		PerPageNums int
		MaxPages    int
		nums        int64
		pageRange   []int
		pageNums    int
		page        int
	}
	type args struct {
		page int
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

			p := &Paginator{
				Request:     tt.fields.Request,
				PerPageNums: tt.fields.PerPageNums,
				MaxPages:    tt.fields.MaxPages,
				nums:        tt.fields.nums,
				pageRange:   tt.fields.pageRange,
				pageNums:    tt.fields.pageNums,
				page:        tt.fields.page,
			}
			if got := p.IsActive(tt.args.page); got != tt.want {
				t.Errorf("Paginator.IsActive() = %v, want %v", got, tt.want)
			}
		})
	}
}

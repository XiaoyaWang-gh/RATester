package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/resources/page"
)

func TestPaginate_1(t *testing.T) {
	type args struct {
		seq     any
		options []any
	}
	tests := []struct {
		name    string
		p       *pagePaginator
		args    args
		want    *page.Pager
		wantErr bool
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

			got, err := tt.p.Paginate(tt.args.seq, tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("pagePaginator.Paginate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pagePaginator.Paginate() = %v, want %v", got, tt.want)
			}
		})
	}
}

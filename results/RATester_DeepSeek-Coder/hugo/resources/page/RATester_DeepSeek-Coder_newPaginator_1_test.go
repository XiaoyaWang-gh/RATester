package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewPaginator_1(t *testing.T) {
	type args struct {
		elements   []paginatedElement
		total      int
		size       int
		urlFactory paginationURLFactory
	}
	tests := []struct {
		name    string
		args    args
		want    *Paginator
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

			got, err := newPaginator(tt.args.elements, tt.args.total, tt.args.size, tt.args.urlFactory)
			if (err != nil) != tt.wantErr {
				t.Errorf("newPaginator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPaginator() = %v, want %v", got, tt.want)
			}
		})
	}
}

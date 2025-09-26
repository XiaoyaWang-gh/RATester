package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestnewPaginatorFromPages_1(t *testing.T) {
	type args struct {
		pages      Pages
		size       int
		urlFactory paginationURLFactory
	}
	tests := []struct {
		name    string
		args    args
		want    *Paginator
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				pages:      []Page{},
				size:       10,
				urlFactory: func(int) string { return "" },
			},
			want:    &Paginator{},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				pages:      []Page{},
				size:       0,
				urlFactory: func(int) string { return "" },
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test Case 3",
			args: args{
				pages:      []Page{},
				size:       -1,
				urlFactory: func(int) string { return "" },
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := newPaginatorFromPages(tt.args.pages, tt.args.size, tt.args.urlFactory)
			if (err != nil) != tt.wantErr {
				t.Errorf("newPaginatorFromPages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPaginatorFromPages() = %v, want %v", got, tt.want)
			}
		})
	}
}

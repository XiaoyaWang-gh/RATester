package pagination

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewPaginator_1(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		req  *http.Request
		per  int
		nums interface{}
	}
	tests := []struct {
		name string
		args args
		want *Paginator
	}{
		{
			name: "TestNewPaginator",
			args: args{
				req:  req,
				per:  10,
				nums: 100,
			},
			want: &Paginator{
				Request:     req,
				PerPageNums: 10,
				nums:        100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewPaginator(tt.args.req, tt.args.per, tt.args.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPaginator() = %v, want %v", got, tt.want)
			}
		})
	}
}

package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPaginate_2(t *testing.T) {
	type args struct {
		seq     any
		options []any
	}
	tests := []struct {
		name    string
		p       *nopPage
		args    args
		want    *Pager
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
				t.Errorf("nopPage.Paginate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nopPage.Paginate() = %v, want %v", got, tt.want)
			}
		})
	}
}

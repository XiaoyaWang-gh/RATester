package pagination

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/core/utils/pagination"
	"github.com/beego/beego/v2/server/web/context"
)

func TestSetPaginator_1(t *testing.T) {
	type args struct {
		context *context.Context
		per     int
		nums    int64
	}
	tests := []struct {
		name          string
		args          args
		wantPaginator *pagination.Paginator
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

			gotPaginator := SetPaginator(tt.args.context, tt.args.per, tt.args.nums)
			if !reflect.DeepEqual(gotPaginator, tt.wantPaginator) {
				t.Errorf("SetPaginator() = %v, want %v", gotPaginator, tt.wantPaginator)
			}
		})
	}
}

package mock

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestMockRawWithCtx_1(t *testing.T) {
	type args struct {
		rs orm.RawSeter
	}
	tests := []struct {
		name string
		args args
		want *Mock
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

			if got := MockRawWithCtx(tt.args.rs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockRawWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}

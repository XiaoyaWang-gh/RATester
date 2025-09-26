package mock

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestNewMock_1(t *testing.T) {
	type args struct {
		cond Condition
		resp []interface{}
		cb   func(inv *orm.Invocation)
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

			if got := NewMock(tt.args.cond, tt.args.resp, tt.args.cb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMock() = %v, want %v", got, tt.want)
			}
		})
	}
}

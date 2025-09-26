package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMockReadOrCreateWithCtx_1(t *testing.T) {
	type args struct {
		tableName string
		cb        func(data interface{})
		insert    bool
		id        int64
		err       error
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

			if got := MockReadOrCreateWithCtx(tt.args.tableName, tt.args.cb, tt.args.insert, tt.args.id, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockReadOrCreateWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}

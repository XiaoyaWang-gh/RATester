package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMockLoadRelatedWithCtx_1(t *testing.T) {
	type args struct {
		tableName string
		name      string
		rows      int64
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

			if got := MockLoadRelatedWithCtx(tt.args.tableName, tt.args.name, tt.args.rows, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockLoadRelatedWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}

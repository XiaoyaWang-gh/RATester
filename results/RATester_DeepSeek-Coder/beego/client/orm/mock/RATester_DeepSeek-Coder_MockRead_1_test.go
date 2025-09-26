package mock

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMockRead_1(t *testing.T) {
	type args struct {
		tableName string
		cb        func(data interface{})
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

			if got := MockRead(tt.args.tableName, tt.args.cb, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockRead() = %v, want %v", got, tt.want)
			}
		})
	}
}

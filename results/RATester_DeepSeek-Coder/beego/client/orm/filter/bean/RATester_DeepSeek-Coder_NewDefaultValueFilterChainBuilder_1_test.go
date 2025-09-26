package bean

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/core/bean"
)

func TestNewDefaultValueFilterChainBuilder_1(t *testing.T) {
	type args struct {
		typeAdapters           map[string]bean.TypeAdapter
		includeInsertOrUpdate  bool
		compatibleWithOldStyle bool
	}
	tests := []struct {
		name string
		args args
		want *DefaultValueFilterChainBuilder
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

			if got := NewDefaultValueFilterChainBuilder(tt.args.typeAdapters, tt.args.includeInsertOrUpdate, tt.args.compatibleWithOldStyle); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultValueFilterChainBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

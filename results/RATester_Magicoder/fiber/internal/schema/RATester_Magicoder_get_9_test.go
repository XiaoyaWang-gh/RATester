package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func Testget_9(t *testing.T) {
	type args struct {
		alias string
	}
	tests := []struct {
		name string
		i    *structInfo
		args args
		want *fieldInfo
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

			if got := tt.i.get(tt.args.alias); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() = %v, want %v", got, tt.want)
			}
		})
	}
}

package json

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetData_4(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		c    *JSONConfigContainer
		args args
		want interface{}
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

			if got := tt.c.getData(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONConfigContainer.getData() = %v, want %v", got, tt.want)
			}
		})
	}
}

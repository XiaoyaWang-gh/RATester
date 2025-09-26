package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func Testclone_1(t *testing.T) {
	type fields struct {
		params []condValue
	}
	tests := []struct {
		name   string
		fields fields
		want   *Condition
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

			c := Condition{
				params: tt.fields.params,
			}
			if got := c.clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

package fiber

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetValues_1(t *testing.T) {
	type fields struct {
		values [maxParams]string
	}
	tests := []struct {
		name   string
		fields fields
		want   *[maxParams]string
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

			c := &DefaultCtx{
				values: tt.fields.values,
			}
			if got := c.getValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

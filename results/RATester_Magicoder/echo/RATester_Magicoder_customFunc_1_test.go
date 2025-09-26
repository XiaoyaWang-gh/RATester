package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCustomFunc_1(t *testing.T) {
	tests := []struct {
		name           string
		sourceParam    string
		customFunc     func(values []string) []error
		valueMustExist bool
		want           *ValueBinder
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

			b := &ValueBinder{}
			if got := b.customFunc(tt.sourceParam, tt.customFunc, tt.valueMustExist); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("customFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

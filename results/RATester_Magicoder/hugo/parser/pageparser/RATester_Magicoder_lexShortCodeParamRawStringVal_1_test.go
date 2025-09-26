package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestlexShortCodeParamRawStringVal_1(t *testing.T) {
	tests := []struct {
		name string
		l    *pageLexer
		typ  ItemType
		want stateFunc
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

			if got := lexShortCodeParamRawStringVal(tt.l, tt.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexShortCodeParamRawStringVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

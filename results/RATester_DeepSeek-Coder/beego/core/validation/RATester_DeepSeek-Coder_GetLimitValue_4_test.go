package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetLimitValue_4(t *testing.T) {
	type fields struct {
		Match Match
		Key   string
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
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

			n := NoMatch{
				Match: tt.fields.Match,
				Key:   tt.fields.Key,
			}
			if got := n.GetLimitValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoMatch.GetLimitValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

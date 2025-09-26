package try

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHasHeaderValue_1(t *testing.T) {
	type args struct {
		header     string
		value      string
		exactMatch bool
	}
	tests := []struct {
		name string
		args args
		want ResponseCondition
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

			if got := HasHeaderValue(tt.args.header, tt.args.value, tt.args.exactMatch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasHeaderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

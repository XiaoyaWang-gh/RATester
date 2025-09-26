package i18n

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetPluralCount_1(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want any
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

			if got := getPluralCount(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPluralCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

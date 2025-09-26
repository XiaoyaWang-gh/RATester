package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReplace_4(t *testing.T) {
	type fields struct {
		gomod *goModule
		owner Module
	}
	tests := []struct {
		name   string
		fields fields
		want   Module
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

			m := &moduleAdapter{
				gomod: tt.fields.gomod,
				owner: tt.fields.owner,
			}
			if got := m.Replace(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moduleAdapter.Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}

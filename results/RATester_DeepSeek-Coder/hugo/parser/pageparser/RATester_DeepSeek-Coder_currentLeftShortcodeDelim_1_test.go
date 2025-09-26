package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCurrentLeftShortcodeDelim_1(t *testing.T) {
	tests := []struct {
		name string
		l    *pageLexer
		want []byte
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

			if got := tt.l.currentLeftShortcodeDelim(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pageLexer.currentLeftShortcodeDelim() = %v, want %v", got, tt.want)
			}
		})
	}
}

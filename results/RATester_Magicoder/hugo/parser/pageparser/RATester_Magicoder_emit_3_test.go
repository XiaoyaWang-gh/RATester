package pageparser

import (
	"fmt"
	"reflect"
	"testing"
)

func Testemit_3(t *testing.T) {
	tests := []struct {
		name string
		l    *pageLexer
		t    ItemType
		want *pageLexer
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

			tt.l.emit(tt.t)
			if !reflect.DeepEqual(tt.l, tt.want) {
				t.Errorf("emit() = %v, want %v", tt.l, tt.want)
			}
		})
	}
}

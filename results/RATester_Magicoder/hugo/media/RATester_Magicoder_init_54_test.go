package media

import (
	"fmt"
	"reflect"
	"testing"
)

func Testinit_54(t *testing.T) {
	type args struct {
		m *Type
	}
	tests := []struct {
		name string
		args args
		want *Type
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

			tt.args.m.init()
			if !reflect.DeepEqual(tt.args.m, tt.want) {
				t.Errorf("init() = %v, want %v", tt.args.m, tt.want)
			}
		})
	}
}

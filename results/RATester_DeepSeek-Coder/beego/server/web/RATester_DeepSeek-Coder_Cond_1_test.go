package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCond_1(t *testing.T) {
	type fields struct {
		prefix   string
		handlers *ControllerRegister
	}
	type args struct {
		cond namespaceCond
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Namespace
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

			n := &Namespace{
				prefix:   tt.fields.prefix,
				handlers: tt.fields.handlers,
			}
			if got := n.Cond(tt.args.cond); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.Cond() = %v, want %v", got, tt.want)
			}
		})
	}
}

package accesslog

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewFieldHandler_1(t *testing.T) {
	type args struct {
		next    http.Handler
		name    string
		value   string
		applyFn FieldApply
	}
	tests := []struct {
		name string
		args args
		want http.Handler
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

			if got := NewFieldHandler(tt.args.next, tt.args.name, tt.args.value, tt.args.applyFn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFieldHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

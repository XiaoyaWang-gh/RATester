package customerrors

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestNewCodeModifier_1(t *testing.T) {
	type args struct {
		rw   http.ResponseWriter
		code int
	}
	tests := []struct {
		name string
		args args
		want *codeModifier
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

			if got := newCodeModifier(tt.args.rw, tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCodeModifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

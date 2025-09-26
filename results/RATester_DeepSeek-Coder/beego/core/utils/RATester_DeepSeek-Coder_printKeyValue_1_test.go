package utils

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestPrintKeyValue_1(t *testing.T) {
	type args struct {
		val          reflect.Value
		pointers     **pointerInfo
		interfaces   *[]reflect.Value
		structFilter func(string, string) bool
		formatOutput bool
		indent       string
		level        int
	}
	tests := []struct {
		name string
		args args
		want string
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

			w := &bytes.Buffer{}
			printKeyValue(w, tt.args.val, tt.args.pointers, tt.args.interfaces, tt.args.structFilter, tt.args.formatOutput, tt.args.indent, tt.args.level)
			if got := w.String(); got != tt.want {
				t.Errorf("printKeyValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

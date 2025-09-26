package utils

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestPrintKeyValue_1(t *testing.T) {
	type args struct {
		buf          *bytes.Buffer
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

			printKeyValue(tt.args.buf, tt.args.val, tt.args.pointers, tt.args.interfaces, tt.args.structFilter, tt.args.formatOutput, tt.args.indent, tt.args.level)
		})
	}
}

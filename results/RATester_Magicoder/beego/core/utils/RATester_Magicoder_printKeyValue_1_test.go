package utils

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestprintKeyValue_1(t *testing.T) {
	type args struct {
		val reflect.Value
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

			buf := &bytes.Buffer{}
			printKeyValue(buf, tt.args.val, nil, nil, nil, false, "", 0)
			if got := buf.String(); got != tt.want {
				t.Errorf("printKeyValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

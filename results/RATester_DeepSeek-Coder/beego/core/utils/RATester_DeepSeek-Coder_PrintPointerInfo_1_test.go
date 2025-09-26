package utils

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPrintPointerInfo_1(t *testing.T) {
	type args struct {
		buf      *bytes.Buffer
		headlen  int
		pointers *pointerInfo
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

			PrintPointerInfo(tt.args.buf, tt.args.headlen, tt.args.pointers)
			if got := tt.args.buf.String(); got != tt.want {
				t.Errorf("PrintPointerInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

package goldmark

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestsanitizeAnchorNameWithHook_1(t *testing.T) {
	type args struct {
		b      []byte
		idType string
		hook   func(buf *bytes.Buffer)
	}
	tests := []struct {
		name string
		args args
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

			if got := sanitizeAnchorNameWithHook(tt.args.b, tt.args.idType, tt.args.hook); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sanitizeAnchorNameWithHook() = %v, want %v", got, tt.want)
			}
		})
	}
}

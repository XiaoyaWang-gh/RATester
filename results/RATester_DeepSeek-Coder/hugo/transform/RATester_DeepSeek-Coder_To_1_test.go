package transform

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestTo_1(t *testing.T) {
	type fields struct {
		from *bytes.Buffer
		to   *bytes.Buffer
	}
	tests := []struct {
		name   string
		fields fields
		want   io.Writer
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

			ft := fromToBuffer{
				from: tt.fields.from,
				to:   tt.fields.to,
			}
			if got := ft.To(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fromToBuffer.To() = %v, want %v", got, tt.want)
			}
		})
	}
}
